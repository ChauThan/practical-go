package alerts

import (
	"context"
	"fmt"
	"math"
	"time"

	"crypto-monitor/internal/binance"
	"crypto-monitor/internal/cache"
)

// Alert represents a significant price change.
type Alert struct {
	Symbol      string    `json:"symbol"`
	OldPrice    float64   `json:"old_price"`
	NewPrice    float64   `json:"new_price"`
	Change      float64   `json:"change"`
	ChangePct   float64   `json:"change_pct"`
	OccurredAt  time.Time `json:"occurred_at"`
	Observation time.Time `json:"observation"`
}

// Engine detects significant price changes.
type Engine struct {
	cache     cache.Cache
	threshold float64
}

// NewEngine builds an alerting engine.
func NewEngine(cache cache.Cache, threshold float64) *Engine {
	return &Engine{cache: cache, threshold: threshold}
}

// Start processes price events and emits alerts.
func (e *Engine) Start(ctx context.Context, prices <-chan binance.PriceEvent) (<-chan Alert, <-chan error) {
	alerts := make(chan Alert, 32)
	errCh := make(chan error, 1)

	go func() {
		defer close(alerts)
		defer close(errCh)

		for {
			select {
			case <-ctx.Done():
				return
			case priceEvent, ok := <-prices:
				if !ok {
					return
				}

				previous, ok, err := e.cache.GetPrice(ctx, priceEvent.Symbol)
				if err != nil {
					errCh <- fmt.Errorf("alerts: cache read failed: %w", err)
					continue
				}

				snapshot := cache.Snapshot{
					Symbol:    priceEvent.Symbol,
					Price:     priceEvent.Price,
					UpdatedAt: priceEvent.EventTime,
				}
				if err := e.cache.SetPrice(ctx, snapshot); err != nil {
					errCh <- fmt.Errorf("alerts: cache write failed: %w", err)
				}

				if !ok || previous.Price == 0 {
					continue
				}

				change := priceEvent.Price - previous.Price
				changePct := (change / previous.Price) * 100
				if math.Abs(changePct) < e.threshold {
					continue
				}

				alert := Alert{
					Symbol:      priceEvent.Symbol,
					OldPrice:    previous.Price,
					NewPrice:    priceEvent.Price,
					Change:      change,
					ChangePct:   changePct,
					OccurredAt:  time.Now().UTC(),
					Observation: priceEvent.EventTime,
				}

				select {
				case <-ctx.Done():
					return
				case alerts <- alert:
				}
			}
		}
	}()

	return alerts, errCh
}
