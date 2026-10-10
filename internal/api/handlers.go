package api

import (
	"encoding/json"
	"net/http"

	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/domain"
	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/engine"
)

type Handler struct {
	evaluator *engine.Evaluator
	txChan    chan<- domain.Transaction
}

func NewHandler(evaluator *engine.Evaluator, txChan chan<- domain.Transaction) *Handler {
	return &Handler{
		evaluator: evaluator,
		txChan:    txChan,
	}
}

func (h *Handler) AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var tx domain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Non-blocking ingest
	select {
	case h.txChan <- tx:
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(map[string]string{"status": "accepted"})
	default:
		http.Error(w, "Service overloaded", http.StatusServiceUnavailable)
	}
}
