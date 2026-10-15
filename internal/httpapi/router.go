package httpapi

import (
	"github.com/romankretov/goledger/internal/database"
	"log/slog"
	"net/http"
)

func NewRouter(store *database.Storage, logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	handler := NewHandler(store, logger)

	mux.HandleFunc("GET /healthz", handler.HandleHealthz)
	mux.HandleFunc("GET /readyz", handler.HandleReadyz)
	// TODO: Register account creation and lookup routes after implementing their handlers.

	return mux
}
