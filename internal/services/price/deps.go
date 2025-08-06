package price

import (
	"context"
	"crypto-price-service/internal/models"
	"time"
)

type Repository interface {
	Create(ctx context.Context, price *models.Price) error
	CreateMany(ctx context.Context, prices []models.Price) error
	ClosestByCoinID(ctx context.Context, coinID string, timestamp time.Time) (*models.Price, error)
	GetAllPricesForCoinByCoinID(ctx context.Context, coinID string) ([]models.Price, error)
}
