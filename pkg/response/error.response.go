package response

import (
	"time"
)

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This package provides a standardized way to handle error responses in the application.
*/

// ErrorResponse is a struct that represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	CodeReason string `json:"errorReason"`
	CreatedAt  string `json:"createdAt"`
}

// The method to return the error response in controller
func BadRequest(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(BAD_REQUEST),
		Message:    message,
		CodeReason: MSG[BAD_REQUEST],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func Unauthorized(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(UNAUTHORIZED),
		Message:    message,
		CodeReason: MSG[UNAUTHORIZED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func Forbidden(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(FORBIDDEN),
		Message:    message,
		CodeReason: MSG[FORBIDDEN],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func NotFound(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(NOT_FOUND),
		Message:    message,
		CodeReason: MSG[NOT_FOUND],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func MethodNotAllowed(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(METHOD_NOT_ALLOWED),
		Message:    message,
		CodeReason: MSG[METHOD_NOT_ALLOWED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func NotAcceptable(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(NOT_ACCEPTABLE),
		Message:    message,
		CodeReason: MSG[NOT_ACCEPTABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func Conflict(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(CONFLICT),
		Message:    message,
		CodeReason: MSG[CONFLICT],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func UnsupportedMediaType(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(UNSUPPORTED_MEDIA_TYPE),
		Message:    message,
		CodeReason: MSG[UNSUPPORTED_MEDIA_TYPE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func InternalServerError(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(INTERNAL_SERVER_ERROR),
		Message:    message,
		CodeReason: MSG[INTERNAL_SERVER_ERROR],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func ServiceUnavailable(message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(SERVICE_UNAVAILABLE),
		Message:    message,
		CodeReason: MSG[SERVICE_UNAVAILABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}

func AnotherError(statusCode int, message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: int(statusCode),
		Message:    message,
		CodeReason: "UNKNOWN_ERROR",
		CreatedAt:  time.Now().Format(time.RFC3339),
	}
}
