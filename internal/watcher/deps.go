package watcher

import (
	"context"
	"crypto-price-service/internal/dto"
	"crypto-price-service/internal/models"
	"time"
)

type Prices interface {
	CreateMany(ctx context.Context, prices []models.Price) error
	ClosestByCoinID(ctx context.Context, coinID string, timestamp time.Time) (*models.Price, error)
}

type Coins interface {
	BySymbols(ctx context.Context, ids []string) ([]models.Coin, error)
	ListActive(ctx context.Context) ([]models.Coin, error)
}

type Fetcher interface {
	FetchPrices(ctx context.Context, ids []string) (dto.CurrentPrices, error)
}
