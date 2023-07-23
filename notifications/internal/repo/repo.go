package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/opentracing/opentracing-go"
	"route256/libs/client/pg"
	"route256/notifications/internal/models"
	"time"
)

type Repository struct {
	client pg.Client
}

func NewRepo(client pg.Client) *Repository {
	return &Repository{client: client}
}

const (
	tableMessageHistory = "msg_history"
)

func (r *Repository) ListUserHistoryDay(ctx context.Context, userID int64, lastMessageTime *time.Time) ([]models.OrderMessage, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Notifications.LomsRepo.ListUserHistoryDay")
	defer span.Finish()

	span.SetTag("userID", userID)

	dayAgo := time.Now().Add(-24 * time.Hour)
	if lastMessageTime == nil || lastMessageTime.Before(dayAgo) {
		lastMessageTime = &dayAgo
	}

	builder := sq.Select("order_id", "status", "created_at").
		From(tableMessageHistory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": userID}).
		Where(sq.Gt{"created_at": lastMessageTime}).
		OrderBy("created_at DESC")

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := pg.Query{
		Name:     "notifications.ListUserHistoryDay",
		QueryRaw: query,
	}

	var msgs []models.OrderMessage
	if err := r.client.PG().ScanAllContext(ctx, &msgs, q, v...); err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *Repository) GetUserIDByOrderID(ctx context.Context, orderID int64) (int64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Notifications.LomsRepo.GetUserIDByOrderID")
	defer span.Finish()

	span.SetTag("orderID", orderID)

	builder := sq.Select("user_id").
		From(tableMessageHistory).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"order_id": orderID}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := pg.Query{
		Name:     "notifications.GetUserIDByOrderID",
		QueryRaw: query,
	}

	var userID int64
	if err := r.client.PG().ScanOneContext(ctx, &userID, q, v...); err != nil {
		return 0, err
	}

	return userID, nil
}
