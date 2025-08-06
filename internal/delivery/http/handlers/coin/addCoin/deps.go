package addCoin

import (
	"context"
	"crypto-price-service/internal/models"
)

type Coins interface {
	Activate(ctx context.Context, coin *models.Coin) error
}
