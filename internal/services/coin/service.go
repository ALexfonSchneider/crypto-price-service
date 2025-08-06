package coin

import (
	"context"
	apperrors "crypto-price-service/internal/errors"
	"crypto-price-service/internal/models"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Activate(ctx context.Context, coin *models.Coin) error {
	stored, err := s.repo.BySymbol(ctx, coin.Symbol)
	if err != nil {
		return err
	}

	if stored == nil {
		return s.repo.Create(ctx, coin)
	}

	return s.repo.Activate(ctx, stored.ID)
}

func (s *Service) Deactivate(ctx context.Context, symbol string) error {
	coin, err := s.repo.BySymbol(ctx, symbol)
	if err != nil {
		return err
	}

	if coin == nil {
		return apperrors.NewCoinNotExists()
	}

	return s.repo.Deactivate(ctx, coin.ID)
}

func (s *Service) ListActive(ctx context.Context) ([]models.Coin, error) {
	coins, err := s.repo.ListActive(ctx)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func (s *Service) BySymbol(ctx context.Context, symbol string) (*models.Coin, error) {
	coin, err := s.repo.BySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}

	if coin == nil {
		return nil, apperrors.NewCoinNotExists()
	}

	return coin, nil
}

func (s *Service) BySymbols(ctx context.Context, symbols []string) ([]models.Coin, error) {
	coins, err := s.repo.BySymbols(ctx, symbols)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
