FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o fraud-engine cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/fraud-engine .

EXPOSE 8080
CMD ["./fraud-engine"]
