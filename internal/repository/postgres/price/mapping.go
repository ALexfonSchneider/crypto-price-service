package price

import (
	db "crypto-price-service/internal/db/gen"
	"crypto-price-service/internal/models"
)

func mapFromDTO(dto db.Prices) models.Price {
	return models.Price{
		ID:        dto.ID,
		CoinID:    dto.CoinID,
		Price:     dto.Price,
		CreatedAt: dto.CreatedAt,
	}
}
