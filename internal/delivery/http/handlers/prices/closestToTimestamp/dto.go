package closestToTimestamp

import "time"

type SuccessResponse struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
