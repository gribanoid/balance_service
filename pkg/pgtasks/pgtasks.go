package pgtasks

import (
	"context"
	"github.com/gribanoid/balance_service/internal/repositories/postgres"

	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
	Tx      pgx.Tx
	conn    *pgxpool.Conn
	timeout time.Duration
}

func (t *Task) Commit(ctx context.Context) error {
	trCtx, cancel := context.WithTimeout(ctx, t.timeout)
	defer cancel()

	if err := t.Tx.Commit(trCtx); err != nil {
		return fmt.Errorf("transaction commit failed: %s", err)
	}

	if t.conn != nil {
		t.conn.Release()
	}

	return nil
}

func (t *Task) Rollback(ctx context.Context) {
	trCtx, cancel := context.WithTimeout(ctx, t.timeout)
	defer cancel()

	if err := t.Tx.Rollback(trCtx); err != nil {
		if err.Error() != "tx is closed" {
			log.Fatal(err)
		}
	}

	if t.conn != nil {
		t.conn.Release()
	}
}

type AbstactStorageTx interface {
	CreateTx(ctx context.Context, t time.Duration) (task *Task, err error)
}

type storageImpl struct {
	db      *postgres.PGConnPool
	timeout time.Duration
}

func NewStorage(db *postgres.PGConnPool, timeout time.Duration) AbstactStorageTx {
	return &storageImpl{db, timeout}
}

func (s *storageImpl) CreateTx(ctx context.Context, t time.Duration) (task *Task, err error) {
	trCtx, cancel := context.WithTimeout(ctx, t)
	defer cancel()

	conn, err := s.db.Acquire(trCtx)
	if err != nil {
		return nil, err
	}

	tx, err := conn.BeginTx(trCtx, pgx.TxOptions{AccessMode: pgx.ReadWrite})
	if err != nil {
		return nil, fmt.Errorf("sql.BeginTx err: %s", err)
	}
	return &Task{
		Tx:      tx,
		conn:    conn,
		timeout: s.timeout,
	}, nil
}
