package errors

var (
	internalServerError = 0
	coinNotExists       = 1000
	InvalidRequest      = 1001
	priceNotFound       = 1002
)

func NewInternalServerError() *ApplicationError {
	return &ApplicationError{
		code:     internalServerError,
		httpCode: 500,
		message:  "Internal server error",
	}
}

func NewCoinNotExists() *ApplicationError {
	return &ApplicationError{
		code:     coinNotExists,
		httpCode: 404,
		message:  "Coin does not exists",
	}
}

func NewInvalidRequest() *ApplicationError {
	return &ApplicationError{
		code:     InvalidRequest,
		httpCode: 400,
		message:  "Invalid request",
	}
}

func NewPriceNotFound() *ApplicationError {
	return &ApplicationError{
		code:     priceNotFound,
		httpCode: 404,
		message:  "Price not found",
	}
}
