package addCoin

type Request struct {
	Name   string `json:"name" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
