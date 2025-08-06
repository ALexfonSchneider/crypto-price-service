package price

import (
	"context"
	db "crypto-price-service/internal/db/gen"
	"crypto-price-service/internal/models"
	"database/sql"
	"errors"
	"time"
)

type Repository struct {
	db db.Querier
}

func New(db db.Querier) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, price *models.Price) error {
	return r.db.CreatePrice(ctx, db.CreatePriceParams{
		CoinID:    price.CoinID,
		Price:     price.Price,
		CreatedAt: price.CreatedAt,
	})
}

func (r *Repository) CreateMany(ctx context.Context, prices []models.Price) error {
	var args []db.CreatePricesParams
	for _, price := range prices {
		args = append(args, db.CreatePricesParams{
			CoinID:    price.CoinID,
			Price:     price.Price,
			CreatedAt: price.CreatedAt,
		})
	}

	_, err := r.db.CreatePrices(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ClosestByCoinID(ctx context.Context, coinID string, createdAt time.Time) (*models.Price, error) {
	result, err := r.db.GetClosestPriceByCoinID(ctx, db.GetClosestPriceByCoinIDParams{
		CoinID:    coinID,
		CreatedAt: createdAt,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	price := mapFromDTO(result)

	return &price, nil
}

func (r *Repository) GetAllPricesForCoinByCoinID(ctx context.Context, coinID string) ([]models.Price, error) {
	result, err := r.db.GetAllPricesForCoinByCoinID(ctx, coinID)
	if err != nil {
		return nil, err
	}

	prices := make([]models.Price, len(result))
	for i, price := range result {
		prices[i] = mapFromDTO(price)
	}

	return prices, nil
}
