package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	pingContext, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := conn.PingContext(pingContext); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
