package cart

import (
	"context"
	"errors"
	"fmt"

	"route256/checkout/internal/models"
	"route256/libs/client/pg"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

const tableName = "items"

type repository struct {
	client pg.Client
}

func NewRepo(client pg.Client) *repository {
	return &repository{client: client}
}

func (r *repository) AddToCart(ctx context.Context, user int64, item *models.ItemData) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("user_id", "sku", "count").
		Suffix("ON CONFLICT (user_id, sku) DO UPDATE SET count = items.count + ?", item.Count).
		Values(user, item.SKU, item.Count)

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "checkout.AddToCart",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCount(ctx context.Context, user int64, sku uint32) (uint16, error) {
	builder := sq.Select("count").From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": user, "sku": sku}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "checkout.GetCount",
		QueryRaw: query,
	}

	var count uint16
	if err = r.client.PG().ScanOneContext(ctx, &count, q, v...); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		fmt.Println(err)

		return 0, err
	}

	return count, nil
}

func (r *repository) GetUserCart(ctx context.Context, user int64) ([]models.ItemData, error) {
	builder := sq.Select("sku", "count").From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": user})

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "checkout.GetUserCart",
		QueryRaw: query,
	}

	var items []models.ItemData
	if err = r.client.PG().ScanAllContext(ctx, &items, q, v...); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		fmt.Println(err)

		return nil, err
	}

	return items, nil
}

func (r *repository) DeleteFromCart(ctx context.Context, user int64, item *models.ItemData) error {
	if err := r.client.RunRepeatableRead(ctx, func(ctx context.Context) error {
		count, err := r.GetCount(ctx, user, item.SKU)
		if err != nil {
			return err
		}

		if count < item.Count {
			ErrStockInsufficient := errors.New("stock insufficient")
			return ErrStockInsufficient
		}

		if count > item.Count {
			return r.removeItemsFromCart(ctx, user, count, item)
		}

		return r.deleteItemFromCart(ctx, user, item)
	}); err != nil {
		return err
	}

	return nil
}

func (r *repository) deleteItemFromCart(ctx context.Context, user int64, item *models.ItemData) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": user, "sku": item.SKU})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "checkout.DeleteFromCart",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) removeItemsFromCart(ctx context.Context, user int64, count uint16, item *models.ItemData) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("count", count-item.Count).
		Where(sq.Eq{"user_id": user, "sku": item.SKU})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "checkout.removeItemsFromCart",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteUserCart(ctx context.Context, user int64) error  {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": user})

	query, v, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := pg.Query{
		Name:     "checkout.DeleteUserCart",
		QueryRaw: query,
	}

	if _, err = r.client.PG().ExecContext(ctx, q, v...); err != nil {
		return err
	}

	return nil
}
