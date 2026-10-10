# Real-Time Fraud Detection Engine

## The Problem
Modern financial systems require ultra-low-latency processing to detect and prevent fraudulent transactions before they settle. Traditional batch-processing pipelines are too slow, and generic rule engines introduce unacceptable overhead. There is a need for a highly concurrent, memory-optimized engine capable of evaluating complex heuristics across thousands of transactions per second while maintaining strict persistence for flagged anomalies.

## The Architecture
This system is architected as a high-throughput, concurrent event processor written in Go. It utilizes Go's goroutines and channels to implement a worker-pool pattern, decoupling HTTP ingestion from rule evaluation. Ingested transactions are routed to an in-memory evaluation engine where deterministic rules are applied. Flagged transactions are asynchronously persisted to a PostgreSQL database using an optimized connection pool, ensuring the critical ingestion path remains unblocked. The entire stack is containerized for seamless deployment.

## Tech Stack
* Programming Language: Go (Golang)
* Database: PostgreSQL
* Infrastructure: Docker, Docker Compose
* API Layer: standard `net/http`

## Local Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/kholipha-ahmmad-al-amin/fraud-detection-engine.git
   cd fraud-detection-engine
   ```
2. Start the infrastructure (PostgreSQL):
   ```bash
   docker-compose up -d db
   ```
3. Run the engine:
   ```bash
   go run cmd/server/main.go
   ```
4. Send a test transaction:
   ```bash
   curl -X POST http://localhost:8080/analyze -d '{"id": "tx-123", "amount": 15000, "user_id": "usr-99", "currency": "USD"}'
   ```

## System Documentation

### System Architecture Diagram
```mermaid
graph TD
    Client[Client Applications] --> |POST /analyze| API[HTTP Ingestion Layer]
    API --> |Transaction| Channel[Buffered Channel]
    Channel --> |Pop| Worker1[Worker Pool]
    Channel --> |Pop| Worker2[Worker Pool]
    Worker1 --> |Evaluate| Engine[Rule Engine]
    Worker2 --> |Evaluate| Engine
    Engine --> |Flagged| DB[(PostgreSQL)]
```

### Entity Relationship Diagram (ERD)
```mermaid
erDiagram
    TRANSACTION {
        string id PK
        string user_id
        float amount
        string currency
        timestamp created_at
    }
    FLAG {
        string id PK
        string transaction_id FK
        string reason
        timestamp flagged_at
    }
    TRANSACTION ||--o{ FLAG : generates
```

### Data Flow Diagram (Level 0)
```mermaid
graph TD
    Client((Client)) --> |Transaction Request| System[Fraud Detection System]
    System --> |Persist Flag| DB[(Database)]
    System --> |Result Response| Client
```

### Data Flow Diagram (Level 1)
```mermaid
graph TD
    Client((Client)) --> |Raw JSON| Ingest(Ingestion API)
    Ingest --> |Parsed Struct| Queue(In-Memory Queue)
    Queue --> |Dispatch| Eval(Evaluation Engine)
    Eval --> |Is Fraud?| RuleCheck{Rules}
    RuleCheck -->|Yes| Store(Persistence Layer)
    Store --> |SQL INSERT| DB[(PostgreSQL)]
    RuleCheck -->|No| Discard(Discard)
```

### Use Case Diagram
```mermaid
flowchart LR
    Client((Client)) --> SubmitTx(Submit Transaction)
    SubmitTx -.-> |include| EvalRules(Evaluate Rules)
    EvalRules -.-> |extend| FlagTx(Flag Transaction)
    Admin((Administrator)) --> ReviewFlags(Review Flags)
```

### Sequence Diagram
```mermaid
sequenceDiagram
    participant C as Client
    participant A as API
    participant W as Worker Pool
    participant R as Rule Engine
    participant D as Database

    C->>A: POST /analyze
    A->>W: Enqueue Transaction
    A-->>C: 202 Accepted
    W->>R: Evaluate(Transaction)
    alt is Fraudulent
        R->>D: Insert Flag Record
    end
```

Developed by Kholipha Ahmmad Al-Amin
Founder and CEO, EquiSaaS BD | Principal Consultant, AR IT Consultancy
Expertise: Full Stack Enterprise Architecture & AI Engineering
GitHub: https://github.com/kholipha-ahmmad-al-amin
LinkedIn: https://www.linkedin.com/in/kholipha-ahmmad-al-amin
Portfolio: https://kholipha-ahmmad-al-amin.equisaas-bd.com
Email: kholifaahmadalamin@gmail.com
