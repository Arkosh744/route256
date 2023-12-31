package cart

import (
	"context"
	"errors"

	"route256/checkout/internal/models"
	"route256/libs/client/pg"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
)

const tableName = "items"

type Repository struct {
	client pg.Client
}

func NewRepo(client pg.Client) *Repository {
	return &Repository{client: client}
}

func (r *Repository) AddToCart(ctx context.Context, user int64, item *models.ItemData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.AddToCart")
	defer span.Finish()

	span.SetTag("userID", user)
	span.SetTag("SKU", item.SKU)
	span.SetTag("count", item.Count)

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

func (r *Repository) GetCount(ctx context.Context, user int64, sku uint32) (uint16, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.GetCount")
	defer span.Finish()

	span.SetTag("userID", user)
	span.SetTag("SKU", sku)

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
		return 0, err
	}

	return count, nil
}

func (r *Repository) GetUserCart(ctx context.Context, user int64) ([]models.ItemData, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.GetUserCart")
	defer span.Finish()

	span.SetTag("userID", user)

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
		return nil, err
	}

	return items, nil
}

func (r *Repository) DeleteFromCart(ctx context.Context, user int64, item *models.ItemData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.DeleteFromCart")
	defer span.Finish()

	span.SetTag("userID", user)
	span.SetTag("SKU", item.SKU)
	span.SetTag("count", item.Count)

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

func (r *Repository) deleteItemFromCart(ctx context.Context, user int64, item *models.ItemData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.deleteItemFromCart")
	defer span.Finish()

	span.SetTag("userID", user)
	span.SetTag("SKU", item.SKU)
	span.SetTag("count", item.Count)

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

func (r *Repository) removeItemsFromCart(ctx context.Context, user int64, count uint16, item *models.ItemData) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.removeItemsFromCart")
	defer span.Finish()

	span.SetTag("userID", user)
	span.SetTag("SKU", item.SKU)
	span.SetTag("count", item.Count)
	span.SetTag("NewCount", count-item.Count)

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

func (r *Repository) DeleteUserCart(ctx context.Context, user int64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Checkout.Repo.DeleteUserCart")
	defer span.Finish()

	span.SetTag("userID", user)

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
