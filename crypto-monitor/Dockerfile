FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal

RUN go build -o /out/crypto-monitor ./cmd

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/crypto-monitor /crypto-monitor

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/crypto-monitor"]
