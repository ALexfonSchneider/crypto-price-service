package coin

import (
	"context"
	"crypto-price-service/internal/models"
)

type Repository interface {
	Create(ctx context.Context, coin *models.Coin) error
	BySymbols(ctx context.Context, ids []string) ([]models.Coin, error)
	BySymbol(ctx context.Context, symbol string) (*models.Coin, error)
	Deactivate(ctx context.Context, id string) error
	Activate(ctx context.Context, id string) error
	ListActive(ctx context.Context) ([]models.Coin, error)
}
