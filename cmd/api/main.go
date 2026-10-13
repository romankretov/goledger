package main

import (
	"context"
	"github.com/romankretov/goledger/internal/database"
	"github.com/romankretov/goledger/internal/httpapi"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	connString := "postgres://goledger:123456@127.0.0.1:5432/goledger?sslmode=disable"

	db, err := database.Connect(ctx, connString)
	if err != nil {
		panic(err)
	}
	cancel()
	store := database.NewStorage(db)
	mux := httpapi.NewRouter(store)

	server := &http.Server{
		Addr:    ":5123",
		Handler: mux,
	}

	server.ListenAndServe()
}
