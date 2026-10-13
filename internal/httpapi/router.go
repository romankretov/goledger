package httpapi

import (
	"github.com/romankretov/goledger/internal/database"
	"net/http"
)

func NewRouter(store *database.Storage) *http.ServeMux {
	mux := http.NewServeMux()
	handler := NewHandler(store)

	mux.HandleFunc("GET /healthz", handler.HandleHealthz)

	return mux
}
