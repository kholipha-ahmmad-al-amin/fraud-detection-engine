package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/api"
	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/domain"
	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/engine"
	"github.com/kholipha-ahmmad-al-amin/fraud-detection-engine/internal/storage"
)

func main() {
	var store *storage.PostgresStore
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		var err error
		store, err = storage.NewPostgresStore(dbURL)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		log.Println("Connected to PostgreSQL")
	} else {
		log.Println("No DATABASE_URL provided. Running without persistence.")
	}

	// Buffer channel for high-throughput ingestion
	txChan := make(chan domain.Transaction, 1000)
	evaluator := engine.NewEvaluator()

	// Initialize Worker Pool
	for i := 0; i < 10; i++ {
		go func(workerID int) {
			for tx := range txChan {
				flag := evaluator.Evaluate(tx)
				if flag != nil {
					log.Printf("[Worker %d] FRAUD DETECTED: %s for TxID: %s\n", workerID, flag.Reason, tx.ID)
					if store != nil {
						store.SaveFlag(flag)
					}
				}
			}
		}(i)
	}

	handler := api.NewHandler(evaluator, txChan)
	mux := http.NewServeMux()
	mux.HandleFunc("/analyze", handler.AnalyzeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Fraud Detection Engine on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
