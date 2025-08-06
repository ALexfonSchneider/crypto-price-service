package dto

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Cause   string `json:"cause,omitempty"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
