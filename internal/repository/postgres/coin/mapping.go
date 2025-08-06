package coin

import (
	db "crypto-price-service/internal/db/gen"
	"crypto-price-service/internal/models"
)

func mapFromDTO(dto db.Coins) models.Coin {
	return models.Coin{
		ID:        dto.ID,
		Name:      dto.Name,
		Symbol:    dto.Symbol,
		IsActive:  dto.IsActive,
		CreatedAt: dto.CreatedAt,
	}
}
