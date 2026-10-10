CREATE TABLE flags (
    id VARCHAR(36) PRIMARY KEY,
    transaction_id VARCHAR(255) NOT NULL,
    reason TEXT NOT NULL,
    flagged_at TIMESTAMP NOT NULL
);
