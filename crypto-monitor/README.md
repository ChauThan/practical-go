# Crypto Monitor

Real-time crypto price monitoring service that streams Binance ticker updates, caches latest prices in Redis, and pushes alert events over an internal WebSocket server.

## 🎯 Learning Objectives

This project demonstrates:
- **Goroutines and channels** for streaming pipelines
- **Context cancellation** for graceful shutdown
- **Dependency injection** via explicit wiring
- **Redis caching** with `go-redis`
- **WebSocket servers/clients** with Gorilla WebSocket

## 🚀 Features

- Binance WebSocket ticker stream ingestion
- Redis-backed latest price cache
- Threshold-based alerting engine
- Internal WebSocket broadcast server for alerts
- Context-driven shutdown and goroutine orchestration
- Multi-stage Docker build (distroless runtime)

## 📁 Project Structure

```
crypto-monitor/
├── cmd/
│   └── main.go
├── internal/
│   ├── alerts/
│   ├── binance/
│   ├── cache/
│   ├── config/
│   └── wsserver/
├── Dockerfile
├── go.mod
└── README.md
```

## ⚙️ Configuration

All configuration is read from environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `BINANCE_SYMBOLS` | `BTCUSDT` | Comma-separated symbol list (e.g. `BTCUSDT,ETHUSDT`). |
| `REDIS_ADDR` | `localhost:6379` | Redis address. |
| `REDIS_PASSWORD` | empty | Redis password. |
| `REDIS_DB` | `0` | Redis database number. |
| `ALERT_THRESHOLD_PCT` | `1.0` | Percent change threshold for alerts. |
| `INTERNAL_WS_ADDR` | `:8080` | WebSocket server bind address. |

## 🛠️ Installation

### Prerequisites

- Go 1.22+
- Redis

### Build

```bash
go build ./...
```

## ▶️ Run

Ensure Redis is running and reachable via `REDIS_ADDR`.

```bash
set REDIS_ADDR=localhost:6379
set BINANCE_SYMBOLS=BTCUSDT,ETHUSDT
set ALERT_THRESHOLD_PCT=0.5
set INTERNAL_WS_ADDR=:8080

go run ./cmd
```

## 📡 WebSocket Alerts

Connect a WebSocket client to:

```
ws://localhost:8080/ws
```

Each message is a JSON-encoded alert:

```json
{
  "symbol": "BTCUSDT",
  "old_price": 65000,
  "new_price": 65550,
  "change": 550,
  "change_pct": 0.846,
  "occurred_at": "2026-02-03T12:00:00Z",
  "observation": "2026-02-03T12:00:00Z"
}
```

## 🐳 Docker

Build the image:

```bash
docker build -t crypto-monitor .
```

Run the container (ensure Redis is reachable from the container):

```bash
docker run --rm -p 8080:8080 \
  -e REDIS_ADDR=host.docker.internal:6379 \
  -e BINANCE_SYMBOLS=BTCUSDT,ETHUSDT \
  -e ALERT_THRESHOLD_PCT=0.5 \
  crypto-monitor
```

## 🐳 Docker Compose

Start the app and Redis together:

```bash
docker compose up --build
```

Connect to the WebSocket:

```bash
npx wscat -c ws://localhost:8080/ws
```

Stop services:

```bash
docker compose down
```

Adjust `ALERT_THRESHOLD_PCT` or `BINANCE_SYMBOLS` in `docker-compose.yml` if alerts are too frequent or too slow.

## 🧪 Testing

```bash
go test ./...
```

## 🔗 Workspace

- [Back to Practical Go workspace README](../README.md)

## Troubleshooting

- If you see Redis connection errors, confirm `REDIS_ADDR` and that Redis is reachable.
- If there are no alerts, lower `ALERT_THRESHOLD_PCT` or confirm symbols in `BINANCE_SYMBOLS`.
- If the WebSocket port is unavailable, change `INTERNAL_WS_ADDR`.
