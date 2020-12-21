package dto

type ErrorResponseDTO struct {
	HttpStatus string `json:"http_status"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func NewErrorResponseDTO(httpStatus string, code int, message string) ErrorResponseDTO {
	return ErrorResponseDTO{
		HttpStatus: httpStatus,
		Code: code,
		Message: message,
	}
}
