package closestToTimestamp

import "time"

type Request struct {
	Symbol    string `json:"symbol"`
	Timestamp int64  `json:"timestamp"`
}

type SuccessResponse struct {
	Symbol    string    `json:"symbol"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
