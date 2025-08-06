package addCoin

import (
	"crypto-price-service/internal/models"
	"time"
)

type Request struct {
	Name   string `json:"name" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`
}

type SuccessResponse struct {
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func responseFromModel(coin models.Coin) SuccessResponse {
	return SuccessResponse{
		Symbol:    coin.Symbol,
		Name:      coin.Name,
		CreatedAt: coin.CreatedAt,
	}
}
