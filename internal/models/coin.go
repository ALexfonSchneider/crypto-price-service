package models

import (
	"github.com/google/uuid"
	"time"
)

type Coin struct {
	ID        string
	Name      string
	Symbol    string
	IsActive  bool
	CreatedAt time.Time
}

func NewCoin(name string, symbol string, isActive bool, createdAt time.Time) *Coin {
	return &Coin{
		ID:        uuid.NewString(),
		Name:      name,
		Symbol:    symbol,
		IsActive:  isActive,
		CreatedAt: createdAt,
	}
}
