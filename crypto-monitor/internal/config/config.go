package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds runtime configuration.
type Config struct {
	Symbols           []string
	RedisAddr         string
	RedisPassword     string
	RedisDB           int
	AlertThresholdPct float64
	InternalWsAddr    string
}

// Load reads configuration from environment variables.
func Load() (Config, error) {
	config := Config{
		Symbols:           []string{"BTCUSDT"},
		RedisAddr:         "localhost:6379",
		RedisPassword:     "",
		RedisDB:           0,
		AlertThresholdPct: 1.0,
		InternalWsAddr:    ":8080",
	}

	if value := strings.TrimSpace(os.Getenv("BINANCE_SYMBOLS")); value != "" {
		config.Symbols = splitCSV(value)
	}
	if value := strings.TrimSpace(os.Getenv("REDIS_ADDR")); value != "" {
		config.RedisAddr = value
	}
	if value := os.Getenv("REDIS_PASSWORD"); value != "" {
		config.RedisPassword = value
	}
	if value := strings.TrimSpace(os.Getenv("REDIS_DB")); value != "" {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return Config{}, fmt.Errorf("invalid REDIS_DB: %w", err)
		}
		config.RedisDB = parsed
	}
	if value := strings.TrimSpace(os.Getenv("ALERT_THRESHOLD_PCT")); value != "" {
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return Config{}, fmt.Errorf("invalid ALERT_THRESHOLD_PCT: %w", err)
		}
		config.AlertThresholdPct = parsed
	}
	if value := strings.TrimSpace(os.Getenv("INTERNAL_WS_ADDR")); value != "" {
		config.InternalWsAddr = value
	}
	return config, nil
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
