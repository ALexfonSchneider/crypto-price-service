package closestToTimestamp

import (
	"context"
	"crypto-price-service/internal/models"
	"time"
)

type Prices interface {
	ClosestByCoinID(ctx context.Context, coinID string, timestamp time.Time) (*models.Price, error)
}
type Coins interface {
	BySymbol(ctx context.Context, symbol string) (*models.Coin, error)
}
