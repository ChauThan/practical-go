package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const priceKeyPrefix = "price:"

// Snapshot holds the latest price for a symbol.
type Snapshot struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Cache defines the operations used by the alerting pipeline.
type Cache interface {
	GetPrice(ctx context.Context, symbol string) (Snapshot, bool, error)
	SetPrice(ctx context.Context, snapshot Snapshot) error
}

// RedisCache stores latest prices in Redis.
type RedisCache struct {
	client *redis.Client
}

// NewRedisClient builds a Redis client with basic settings.
func NewRedisClient(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// NewRedisCache wraps a Redis client with caching helpers.
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

// GetPrice returns a snapshot if present.
func (c *RedisCache) GetPrice(ctx context.Context, symbol string) (Snapshot, bool, error) {
	key := priceKey(symbol)
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return Snapshot{}, false, nil
		}
		return Snapshot{}, false, fmt.Errorf("redis get failed: %w", err)
	}

	var snapshot Snapshot
	if err := json.Unmarshal([]byte(value), &snapshot); err != nil {
		return Snapshot{}, false, fmt.Errorf("redis decode failed: %w", err)
	}

	return snapshot, true, nil
}

// SetPrice stores the latest snapshot.
func (c *RedisCache) SetPrice(ctx context.Context, snapshot Snapshot) error {
	key := priceKey(snapshot.Symbol)
	payload, err := json.Marshal(snapshot)
	if err != nil {
		return fmt.Errorf("redis encode failed: %w", err)
	}

	if err := c.client.Set(ctx, key, payload, 0).Err(); err != nil {
		return fmt.Errorf("redis set failed: %w", err)
	}

	return nil
}

func priceKey(symbol string) string {
	return priceKeyPrefix + strings.ToUpper(symbol)
}
