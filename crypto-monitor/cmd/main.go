package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"crypto-monitor/internal/alerts"
	"crypto-monitor/internal/binance"
	"crypto-monitor/internal/cache"
	"crypto-monitor/internal/config"
	"crypto-monitor/internal/wsserver"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	redisClient := cache.NewRedisClient(config.RedisAddr, config.RedisPassword, config.RedisDB)
	cacheStore := cache.NewRedisCache(redisClient)
	binanceClient := binance.NewClient()
	alertEngine := alerts.NewEngine(cacheStore, config.AlertThresholdPct)
	wsServer := wsserver.NewServer(config.InternalWsAddr)

	prices, priceErrs := binanceClient.StreamTickers(ctx, config.Symbols)
	alertStream, alertErrs := alertEngine.Start(ctx, prices)

	go wsServer.Broadcast(ctx, alertStream)
	go logErrors(ctx, "binance", priceErrs)
	go logErrors(ctx, "alerts", alertErrs)

	if err := wsServer.Run(ctx); err != nil {
		log.Printf("ws server error: %v", err)
	}

	if err := redisClient.Close(); err != nil {
		log.Printf("redis close error: %v", err)
	}
}

func logErrors(ctx context.Context, source string, errs <-chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		case err, ok := <-errs:
			if !ok {
				return
			}
			log.Printf("%s error: %v", source, err)
		}
	}
}
