package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	// TODO: Wrap database errors with %w so callers can inspect their underlying causes.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingContext, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	if err := db.PingContext(pingContext); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return db, nil
}

func (s *Storage) Ping(ctx context.Context) error {
	err := s.db.PingContext(ctx)
	return err
}

// TODO: Add context-aware account creation and lookup methods, preferably in a separate accounts.go.
// TODO: Define how lookup reports a missing account so handlers can return 404.
// TODO: Add PostgreSQL integration tests for account persistence and database constraints.
// TODO: Before transfers, document money units, balanced immutable entries, overdrafts,
// initial funding, and the source of truth for balances.
// TODO: Implement transfers in one SQL transaction with a consistent account lock order.
// TODO: Add idempotency and test competing withdrawals, reverse transfers, retries, and rollback.
