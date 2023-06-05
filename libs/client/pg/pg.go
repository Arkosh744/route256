package pg

import (
	"context"

	"route256/libs/log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Query struct {
	Name     string
	QueryRaw string
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type PG interface {
	QueryExecer
	NamedExecer
	Pinger

	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close() error
}

type pg struct {
	pgxPool *pgxpool.Pool
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.pgxPool.BeginTx(ctx, txOptions)
}

func (p *pg) Close() error {
	p.pgxPool.Close()

	return nil
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pgxPool.Ping(ctx)
}

func (p *pg) ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error) {
	log.Infof("%s; %s", q.QueryRaw, args)

	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(pgx.Tx).Exec(ctx, q.QueryRaw, args...)
	}

	return p.pgxPool.Exec(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error) {
	log.Infof("%s; %s", q.QueryRaw, args)

	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(pgx.Tx).Query(ctx, q.QueryRaw, args...)
	}

	return p.pgxPool.Query(ctx, q.QueryRaw, args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row {
	log.Infof("%s; %s", q.QueryRaw, args)

	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(pgx.Tx).QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.pgxPool.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}