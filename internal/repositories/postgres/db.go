package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PGConnPool struct {
	pool *pgxpool.Pool
}

func (p *PGConnPool) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	return p.pool.Acquire(ctx)
}

func (p *PGConnPool) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func NewPgxPool(ctx context.Context, connString string, logLevel pgx.LogLevel) (*PGConnPool, error) {

	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conf.LazyConnect = true

	if logLevel != 0 {
		conf.ConnConfig.LogLevel = logLevel
	}

	pool, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, errors.WithMessagef(err, "pgx connection error")
	}

	return &PGConnPool{
		pool: pool,
	}, nil
}
