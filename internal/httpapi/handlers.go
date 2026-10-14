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

func (h *Handler) HandleReadyz(w http.ResponseWriter, r *http.Request) {
	success, err := h.store.Ping(r.Context())
	if !success {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

}
