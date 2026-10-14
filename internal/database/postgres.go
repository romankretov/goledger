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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %v", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingContext, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := db.PingContext(pingContext); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %v", err)
	}

	return db, nil
}

func (s *Storage) Ping(ctx context.Context) (bool, error) {
	err := s.db.PingContext(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}
