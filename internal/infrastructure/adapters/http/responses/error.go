package responses

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BadRequestError(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}
func NotFoundError(message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusNotFound,
		Message: message,
	}
}
