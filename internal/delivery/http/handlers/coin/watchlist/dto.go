package watchlist

import (
	"crypto-price-service/internal/models"
	"time"
)

type Coin struct {
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func fromModel(coins []models.Coin) []Coin {
	result := make([]Coin, len(coins))
	for i, coin := range coins {
		result[i] = Coin{
			Name:      coin.Name,
			Symbol:    coin.Symbol,
			IsActive:  coin.IsActive,
			CreatedAt: coin.CreatedAt,
		}
	}

	return result
}

type SuccessResponse struct {
	Coins []Coin `json:"coins"`
}
