package watcher

import (
	"context"
	"crypto-price-service/internal/models"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type Watcher struct {
	step    time.Duration
	prices  Prices
	coins   Coins
	fetcher Fetcher

	logger *slog.Logger
}

func New(prices Prices, coins Coins, fetcher Fetcher, logger *slog.Logger, step time.Duration) *Watcher {
	watcher := &Watcher{
		prices:  prices,
		coins:   coins,
		fetcher: fetcher,
		step:    step,
	}

	if logger == nil {
		logger = slog.Default()
	}
	watcher.logger = logger

	return watcher
}

func getSymbolToCoinIDMap(coins []models.Coin) map[string]string {
	symbolMap := make(map[string]string)
	for _, coin := range coins {
		symbolMap[coin.Symbol] = coin.ID
	}
	return symbolMap
}

func (w *Watcher) watch(ctx context.Context, logger *slog.Logger) error {
	now := time.Now()

	logger.Info("Starting watch loop", slog.Time("time", now))

	coins, err := w.coins.ListActive(ctx)
	if err != nil {
		return err
	}

	logger.Info("Have active coins", slog.Int("count", len(coins)))

	var symbols []string
	for _, coin := range coins {
		symbols = append(symbols, coin.Symbol)
	}

	currentPrices, err := w.fetcher.FetchPrices(ctx, symbols)
	if err != nil {
		return err
	}

	logger.Info("Got coin prices", slog.Int("count", len(currentPrices)))

	symbolToCoinIDMap := getSymbolToCoinIDMap(coins)

	var prices []models.Price
	for symbol, price := range currentPrices {
		if coinID, ok := symbolToCoinIDMap[symbol]; ok {
			prices = append(prices, *models.NewPrice(coinID, price, now))
		} else {
			// Можно блокировать несуществующие монеты
		}
	}

	if err := w.prices.CreateMany(ctx, prices); err != nil {
		return err
	}

	logger.Info("Current prices saved")

	return nil
}

func (w *Watcher) Watch(ctx context.Context) error {
	w.logger.Info("Start coin watch service")

	ticker := time.NewTicker(w.step)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger := w.logger.With("id", uuid.NewString())

			childCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()

			if err := w.watch(childCtx, logger); err != nil {
				logger.Error("failed to watch", "error", err)
				time.Sleep(w.step) // можно добавить backoff
			}
		case <-ctx.Done():
			w.logger.Info("context canceled. watcher stopped")
			return nil
		}
	}
}
