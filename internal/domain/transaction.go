package domain

import "time"

type Transaction struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Flag struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	Reason        string    `json:"reason"`
	FlaggedAt     time.Time `json:"flagged_at"`
}
