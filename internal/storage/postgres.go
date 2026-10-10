package storage

import (
	"database/sql"
	"log"
	"time"

	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/domain"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) SaveFlag(flag *domain.Flag) error {
	query := `INSERT INTO flags (id, transaction_id, reason, flagged_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, flag.ID, flag.TransactionID, flag.Reason, flag.FlaggedAt)
	if err != nil {
		log.Printf("Failed to insert flag: %v", err)
		return err
	}
	return nil
}
