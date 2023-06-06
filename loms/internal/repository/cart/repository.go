package cart

import (
	"context"
	"errors"

	"route256/libs/client/pg"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
	"go.uber.org/multierr"
)

type repository struct {
	client pg.Client
}

func NewRepo(client pg.Client) *repository {
	return &repository{client: client}
}

const (
	tableItems       = "order_items"
	tableOrder       = "orders"
	tableReservation = "reservations"
	tableStock       = "stocks"
)

func (r *repository) GetStocks(ctx context.Context, sku uint32) ([]models.StockItem, error) {
	builder := sq.Select("warehouse_id", "count").
		From(tableStock).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"sku": sku})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "loms.GetStocks",
		QueryRaw: query,
	}

	var stocks []models.StockItem
	if err = r.client.PG().ScanAllContext(ctx, &stocks, q, v...); err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *repository) CreateOrder(ctx context.Context, user int64, items []models.Item) (int64, error) {
	orderID, err := r.createOrder(ctx, user)
	if err != nil {
		return 0, err
	}

	if err = r.client.RunRepeatableRead(ctx, func(ctx context.Context) error {
		var stocks []models.StockItem
		for _, item := range items {
			stocks, err = r.GetStocks(ctx, item.SKU)
			if err != nil {
				return err
			}

			toReserve := uint64(item.Count)
			for _, stock := range stocks {
				if stock.Count > toReserve {
					if err = r.createReservation(ctx, orderID, stock.WarehouseID, item.SKU, toReserve); err != nil {
						return err
					}

					if err = r.updateStock(ctx, stock.WarehouseID, item.SKU, stock.Count-toReserve); err != nil {
						return err
					}

					toReserve = 0
					break
				}

				if stock.Count == toReserve {
					if err = r.createReservation(ctx, orderID, stock.WarehouseID, item.SKU, toReserve); err != nil {
						return err
					}

					if stock.Count-toReserve == 0 {
						if err = r.deleteStock(ctx, stock.WarehouseID, item.SKU); err != nil {
							return err
						}
					}

					toReserve = 0
					break
				}

				if err = r.createReservation(ctx, orderID, stock.WarehouseID, item.SKU, stock.Count); err != nil {
					return err
				}

				if err = r.updateStock(ctx, stock.WarehouseID, item.SKU, 0); err != nil {
					return err
				}

				toReserve -= stock.Count
			}

			if toReserve > 0 {
				ErrStockInsufficient := errors.New("stock insufficient")
				return ErrStockInsufficient
			}

		}

		if err = r.createOrderItems(ctx, orderID, items); err != nil {
			return err
		}

		if err = r.updateOrderStatus(ctx, orderID, models.OrderStatusAwaitingPayment); err != nil {
			return err
		}

		return nil
	}); err != nil {
		if updErr := r.updateOrderStatus(ctx, orderID, models.OrderStatusFailed); updErr != nil {
			err = multierr.Append(err, errors.New("failed to update order status to 'failed'"))
		}

		return 0, err
	}

	return orderID, nil
}

func (r *repository) createOrder(ctx context.Context, userId int64) (int64, error) {
	builder := sq.Insert(tableOrder).
		Columns("user_id", "status").
		Values(userId, models.OrderStatusNew).
		Suffix("RETURNING order_id").
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "loms.createOrder",
		QueryRaw: query,
	}

	var orderId int64
	if err = r.client.PG().ScanOneContext(ctx, &orderId, q, v...); err != nil {
		return 0, err
	}

	return orderId, nil
}

func (r *repository) createOrderItems(ctx context.Context, orderID int64, items []models.Item) error {
	builder := sq.Insert(tableItems).
		Columns("order_id", "sku", "count").
		PlaceholderFormat(sq.Dollar)

	for i := range items {
		builder = builder.Values(orderID, items[i].SKU, items[i].Count)
	}

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.createOrderItems",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) updateOrderStatus(ctx context.Context, orderID int64, status string) error {
	builder :=
		sq.Update(tableOrder).
			Set("status", status).
			Where(sq.Eq{"order_id": orderID}).
			PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.updateOrderStatus",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) createReservation(ctx context.Context, orderID, warID int64, sku uint32, count uint64) error {
	builder := sq.Insert(tableReservation).
		Columns("order_id", "warehouse_id", "sku", "count").
		Values(orderID, warID, sku, count).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.createReservation",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) updateStock(ctx context.Context, warehouseID int64, sku uint32, count uint64) error {
	builder :=
		sq.Update(tableStock).
			Set("count", count).
			Where(sq.Eq{"warehouse_id": warehouseID, "sku": sku}).
			PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.updateStock",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) deleteStock(ctx context.Context, warehouseID int64, sku uint32) error {
	builder := sq.Delete(tableStock).
		Where(sq.Eq{"warehouse_id": warehouseID, "sku": sku}).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.deleteStock",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}
