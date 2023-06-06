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
	if err = r.client.PG().ScanAllContext(ctx, &stocks, q, v...); err != nil {
		return nil, err
	}

	return stocks, nil
}
