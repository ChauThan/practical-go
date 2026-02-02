package binance

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const defaultBaseURL = "wss://stream.binance.com:9443"

// Client connects to Binance WebSocket feeds.
type Client struct {
	baseURL          string
	dialer           *websocket.Dialer
	reconnectBackoff time.Duration
}

// PriceEvent represents a normalized price update.
type PriceEvent struct {
	Symbol    string
	Price     float64
	RawPrice  string
	EventTime time.Time
}

type combinedStreamMessage struct {
	Stream string          `json:"stream"`
	Data   json.RawMessage `json:"data"`
}

// NewClient creates a Binance WebSocket client with defaults.
func NewClient() *Client {
	return &Client{
		baseURL:          defaultBaseURL,
		dialer:           websocket.DefaultDialer,
		reconnectBackoff: 2 * time.Second,
	}
}

// StreamTickers connects to Binance and emits price events until context cancel.
func (c *Client) StreamTickers(ctx context.Context, symbols []string) (<-chan PriceEvent, <-chan error) {
	out := make(chan PriceEvent, 32)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		if len(symbols) == 0 {
			errCh <- errors.New("binance: at least one symbol required")
			return
		}

		backoff := c.reconnectBackoff
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			conn, _, err := c.dialer.Dial(c.buildStreamURL(symbols), nil)
			if err != nil {
				errCh <- fmt.Errorf("binance: dial failed: %w", err)
				timer := time.NewTimer(backoff)
				select {
				case <-ctx.Done():
					timer.Stop()
					return
				case <-timer.C:
				}
				backoff = minDuration(backoff*2, 30*time.Second)
				continue
			}
			backoff = c.reconnectBackoff

			err = c.readLoop(ctx, conn, out)
			_ = conn.Close()
			if err != nil && !errors.Is(err, context.Canceled) {
				errCh <- err
			}
		}
	}()

	return out, errCh
}

func (c *Client) buildStreamURL(symbols []string) string {
	streams := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		streams = append(streams, fmt.Sprintf("%s@ticker", strings.ToLower(symbol)))
	}

	query := url.Values{}
	query.Set("streams", strings.Join(streams, "/"))

	return fmt.Sprintf("%s/stream?%s", c.baseURL, query.Encode())
}

func (c *Client) readLoop(ctx context.Context, conn *websocket.Conn, out chan<- PriceEvent) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		_, payload, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("binance: read failed: %w", err)
		}

		var msg combinedStreamMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			return fmt.Errorf("binance: decode failed: %w", err)
		}

		var data map[string]json.RawMessage
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			return fmt.Errorf("binance: data decode failed: %w", err)
		}

		symbolRaw, ok := data["s"]
		if !ok {
			return errors.New("binance: missing symbol")
		}
		var symbol string
		if err := json.Unmarshal(symbolRaw, &symbol); err != nil {
			return fmt.Errorf("binance: symbol parse failed: %w", err)
		}

		priceRaw, ok := data["c"]
		if !ok {
			return errors.New("binance: missing close price")
		}
		price, rawPrice, err := parsePrice(priceRaw)
		if err != nil {
			return fmt.Errorf("binance: price parse failed: %w", err)
		}

		eventRaw, ok := data["E"]
		if !ok {
			return errors.New("binance: missing event time")
		}
		eventTimeMs, err := parseEventTime(eventRaw)
		if err != nil {
			return fmt.Errorf("binance: event time parse failed: %w", err)
		}

		event := PriceEvent{
			Symbol:    symbol,
			Price:     price,
			RawPrice:  rawPrice,
			EventTime: time.UnixMilli(eventTimeMs),
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- event:
		}
	}
}

func parseEventTime(raw json.RawMessage) (int64, error) {
	if len(raw) == 0 {
		return 0, errors.New("missing event time")
	}

	var asInt int64
	if err := json.Unmarshal(raw, &asInt); err == nil {
		return asInt, nil
	}

	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		value, err := strconv.ParseInt(asString, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid event time string: %w", err)
		}
		return value, nil
	}

	return 0, fmt.Errorf("unsupported event time payload: %s", string(raw))
}

func parsePrice(raw json.RawMessage) (float64, string, error) {
	if len(raw) == 0 {
		return 0, "", errors.New("missing price")
	}

	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		price, err := strconv.ParseFloat(asString, 64)
		if err != nil {
			return 0, "", fmt.Errorf("invalid price string: %w", err)
		}
		return price, asString, nil
	}

	var asFloat float64
	if err := json.Unmarshal(raw, &asFloat); err == nil {
		return asFloat, strconv.FormatFloat(asFloat, 'f', -1, 64), nil
	}

	return 0, "", fmt.Errorf("unsupported price payload: %s", string(raw))
}

func minDuration(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
