package cart

import (
	"context"

	"route256/libs/client/pg"
	"route256/loms/internal/models"

	sq "github.com/Masterminds/squirrel"
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
	if err := r.client.PG().ScanAllContext(ctx, &stocks, q, v...); err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *repository) GetReservations(ctx context.Context, orderID int64) ([]models.ReservationItem, error) {
	builder := sq.Select("sku", "warehouse_id", "count").
		From(tableReservation).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_id": orderID})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "loms.GetReservation",
		QueryRaw: query,
	}

	var resItems []models.ReservationItem
	if err := r.client.PG().ScanAllContext(ctx, &resItems, q, v...); err != nil {
		return nil, err
	}

	return resItems, nil
}

func (r *repository) GetOrder(ctx context.Context, orderID int64) (*models.Order, error) {
	builder := sq.Select("user_id", "status").
		From(tableOrder).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_id": orderID}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "loms.GetOrder",
		QueryRaw: query,
	}

	var items models.Order
	if err := r.client.PG().ScanOneContext(ctx, &items, q, v...); err != nil {
		return nil, err
	}

	return &items, nil
}

func (r *repository) GetOrderItems(ctx context.Context, orderID int64) ([]models.Item, error) {
	builder := sq.Select("sku", "count").
		From(tableItems).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_id": orderID})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "loms.GetOrderItems",
		QueryRaw: query,
	}

	var items []models.Item
	if err := r.client.PG().ScanAllContext(ctx, &items, q, v...); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository) DeleteReservation(ctx context.Context, orderID int64) error {
	builder := sq.Delete(tableReservation).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_id": orderID})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.DeleteReservation",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateOrder(ctx context.Context, user int64) (int64, error) {
	builder := sq.Insert(tableOrder).
		Columns("user_id", "status").
		Values(user, models.OrderStatusNew).
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
	if err := r.client.PG().ScanOneContext(ctx, &orderId, q, v...); err != nil {
		return 0, err
	}

	return orderId, nil
}

func (r *repository) CreateOrderItems(ctx context.Context, orderID int64, items []models.Item) error {
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
		Name:     "loms.CreateOrderItems",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) InsertStock(ctx context.Context, item models.ReservationItem) error {
	builder := sq.Insert(tableStock).
		Columns("warehouse_id", "sku", "count").
		Values(item.WarehouseID, item.SKU, item.Count).
		Suffix("ON CONFLICT (warehouse_id, sku) DO UPDATE SET count = stocks.count + ?", item.Count).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.InsertStock",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateOrderStatus(ctx context.Context, orderID int64, status string) error {
	builder := sq.Update(tableOrder).
		Set("status", status).
		Where(sq.Eq{"order_id": orderID}).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.UpdateOrderStatus",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateReservation(ctx context.Context, orderID, warID int64, sku uint32, count uint64) error {
	builder := sq.Insert(tableReservation).
		Columns("order_id", "warehouse_id", "sku", "count").
		Values(orderID, warID, sku, count).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.CreateReservation",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateStock(ctx context.Context, warehouseID int64, sku uint32, count uint64) error {
	builder := sq.Update(tableStock).
		Set("count", count).
		Where(sq.Eq{"warehouse_id": warehouseID, "sku": sku}).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.UpdateStock",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteStock(ctx context.Context, warehouseID int64, sku uint32) error {
	builder := sq.Delete(tableStock).
		Where(sq.Eq{"warehouse_id": warehouseID, "sku": sku}).
		PlaceholderFormat(sq.Dollar)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "loms.DeleteStock",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}
