package AllForCoin

import (
	"context"
	"crypto-price-service/internal/models"
)

type Prices interface {
	GetAllPricesForCoinByCoinID(ctx context.Context, coinID string) ([]models.Price, error)
}
type Coins interface {
	BySymbol(ctx context.Context, symbol string) (*models.Coin, error)
}
