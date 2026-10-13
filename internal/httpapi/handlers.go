package httpapi

import (
	"encoding/json"
	"github.com/romankretov/goledger/internal/database"
	"log/slog"
	"net/http"
)

type Handler struct {
	store  *database.Storage
	logger *slog.Logger
}

func NewHandler(store *database.Storage) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		h.logger.Info("GET /healthz", "request", r)
	}
}
