package price

import (
	"context"
	apperrors "crypto-price-service/internal/errors"
	"crypto-price-service/internal/models"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(ctx context.Context, price *models.Price) error {
	return s.repo.Create(ctx, price)
}

func (s *Service) CreateMany(ctx context.Context, prices []models.Price) error {
	return s.repo.CreateMany(ctx, prices)
}

func (s *Service) ClosestByCoinID(ctx context.Context, coinID string, timestamp time.Time) (*models.Price, error) {
	price, err := s.repo.ClosestByCoinID(ctx, coinID, timestamp)
	if err != nil {
		return nil, err
	}

	if price == nil {
		return nil, apperrors.NewPriceNotFound()
	}

	return price, nil
}

func (s *Service) GetAllPricesForCoinByCoinID(ctx context.Context, coinID string) ([]models.Price, error) {
	prices, err := s.repo.GetAllPricesForCoinByCoinID(ctx, coinID)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
