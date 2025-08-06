package AllForCoin

import (
	"crypto-price-service/internal/models"
	"time"
)

type CoinPrice struct {
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type SuccessResponse struct {
	Prices []CoinPrice `json:"prices"`
}

func FromModel(prices []models.Price) SuccessResponse {
	result := make([]CoinPrice, len(prices))
	for i, price := range prices {
		result[i] = CoinPrice{
			Price:     price.Price,
			CreatedAt: price.CreatedAt,
		}
	}

	return SuccessResponse{result}
}
