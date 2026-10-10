package engine

import (
	"time"

	"github.com/google/uuid"
	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/domain"
)

type Evaluator struct{}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Evaluate(tx domain.Transaction) *domain.Flag {
	// Rule 1: High value transaction
	if tx.Amount > 10000 {
		return &domain.Flag{
			ID:            uuid.New().String(),
			TransactionID: tx.ID,
			Reason:        "Amount exceeds $10,000 threshold",
			FlaggedAt:     time.Now(),
		}
	}

	// Rule 2: Simulated banned user list
	if tx.UserID == "usr-banned" {
		return &domain.Flag{
			ID:            uuid.New().String(),
			TransactionID: tx.ID,
			Reason:        "User is on global ban list",
			FlaggedAt:     time.Now(),
		}
	}

	return nil
}
