package main

import (
	"context"
	"github.com/romankretov/goledger/internal/database"
	"github.com/romankretov/goledger/internal/httpapi"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// TODO: Read the database DSN and server address from configuration; validate required values.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connString := "postgres://goledger:123456@127.0.0.1:5432/goledger?sslmode=disable"

	db, err := database.Connect(ctx, connString)
	// TODO: Log startup failures clearly and arrange db.Close after a successful connection.
	if err != nil {
		panic(err)
	}
	defer db.Close()
	store := database.NewStorage(db)
	mux := httpapi.NewRouter(store, logger)

	server := &http.Server{
		// TODO: Set HTTP timeouts, including ReadHeaderTimeout, to bound slow requests.
		Addr:              ":5123",
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 10,
	}

	// TODO: Handle ListenAndServe errors; treat http.ErrServerClosed as a normal shutdown.
	// TODO: Handle termination signals with bounded server.Shutdown before closing the database.
	server.ListenAndServe()
}
