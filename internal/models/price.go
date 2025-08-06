package models

import (
	"github.com/google/uuid"
	"time"
)

type Price struct {
	ID        string
	CoinID    string
	Price     float64
	CreatedAt time.Time
}

func NewPrice(coinID string, price float64, timestamp time.Time) *Price {
	return &Price{
		ID:        uuid.NewString(),
		CoinID:    coinID,
		Price:     price,
		CreatedAt: timestamp,
	}
}
