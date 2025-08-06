// internal/db/coin_repository.go

package coin

import (
	"context"
	db "crypto-price-service/internal/db/gen"
	"crypto-price-service/internal/models"
	"database/sql"
	"errors"
)

type Repository struct {
	db db.Querier
}

func New(db db.Querier) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, coin *models.Coin) error {
	return r.db.CreateCoin(ctx, db.CreateCoinParams{
		ID:        coin.ID,
		Name:      coin.Name,
		Symbol:    coin.Symbol,
		IsActive:  coin.IsActive,
		CreatedAt: coin.CreatedAt,
	})
}

func (r *Repository) Activate(ctx context.Context, id string) error {
	return r.db.ActivateCoin(ctx, id)
}

func (r *Repository) BySymbols(ctx context.Context, symbols []string) ([]models.Coin, error) {
	result, err := r.db.GetCoinsBySymbols(ctx, symbols)
	if err != nil {
		return nil, err
	}

	var coins []models.Coin //coins := make([]models.Coin, 0, len(symbols))
	for _, coin := range result {
		coins = append(coins, mapFromDTO(coin))
	}

	return coins, nil
}

func (r *Repository) BySymbol(ctx context.Context, symbol string) (*models.Coin, error) {
	result, err := r.db.GetCoinBySymbol(ctx, symbol)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	coin := mapFromDTO(result)

	return &coin, nil
}

func (r *Repository) Deactivate(ctx context.Context, id string) error {
	if err := r.db.DeactivateCoin(ctx, id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListActive(ctx context.Context) ([]models.Coin, error) {
	result, err := r.db.ListActiveCoins(ctx)
	if err != nil {
		return nil, err
	}

	coins := make([]models.Coin, len(result))
	for i, coin := range result {
		coins[i] = mapFromDTO(coin)
	}

	return coins, nil
}
