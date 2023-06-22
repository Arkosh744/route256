package pg

import (
	"context"

	"route256/libs/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/multierr"
)

var _ Client = (*client)(nil)

type Client interface {
	TxManager
	PG() PG

	Close() error
}

type TxManager interface {
	RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
}

type client struct {
	pg PG
}

func NewClient(ctx context.Context, pgCfg *pgxpool.Config) (Client, error) {
	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		return nil, err
	}

	log.Info("pg connected successfully")

	return &client{pg: &pg{pgxPool: dbc}}, nil
}

func (c *client) PG() PG {
	return c.pg
}

func (c *client) Close() error {
	if c.pg != nil {
		return c.pg.Close()
	}

	return nil
}

type txKey string

const key = txKey("tx")

func (c *client) RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error {
	tx, err := c.pg.BeginTx(ctx,
		pgx.TxOptions{
			IsoLevel: pgx.RepeatableRead,
		})
	if err != nil {
		return err
	}

	if err = fx(context.WithValue(ctx, key, tx)); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	if err = tx.Commit(ctx); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}
