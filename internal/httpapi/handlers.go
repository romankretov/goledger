package httpapi

import (
	"context"
	"encoding/json"
	"github.com/romankretov/goledger/internal/database"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	store  *database.Storage
	logger *slog.Logger
}

func NewHandler(store *database.Storage, logger *slog.Logger) *Handler {
	return &Handler{
		store:  store,
		logger: logger,
	}
}

func (h *Handler) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		h.logger.Error("failed to write to healthz response",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err)
	}
}

func (h *Handler) HandleReadyz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()
	err := h.store.Ping(ctx)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{"error": "service unavailable"})
		h.logger.Error("postgres is not reachable",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	// TODO: Check readiness response encoding errors and log them without writing another status.

}

// TODO: Add account creation and lookup handlers with validation and consistent error responses.
// TODO: Test health/readiness status codes and JSON headers, including unavailable DB and timeout cases.
