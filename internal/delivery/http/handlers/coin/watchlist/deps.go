package watchlist

import (
	"context"
	"crypto-price-service/internal/models"
)

type Coins interface {
	ListActive(ctx context.Context) ([]models.Coin, error)
}
