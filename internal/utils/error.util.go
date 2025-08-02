package utils

import (
	"auth_service/proto/common"
)

var ErrorMessage = map[string]string{
	"DatabaseError": "Database operation failed",
	"NotFoundError": "Resource not found",
	"RuntimeError":  "An unexpected error occurred",
}

func DatabaseError() *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_DATABASE_ERROR,
		Message: "Database operation failed",
	}
	return e
}

func NotFoundError() *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_NOT_FOUND,
		Message: "Not found error",
	}
	return e
}

func RuntimeError() *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_RUN_TIME_ERROR,
		Message: "An unexpected error occurred: runtime error",
	}
	return e
}

func UnauthorizedError() *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_UNAUTHORIZED,
		Message: "Unauthorized access",
	}
	return e
}

func PermissionDeniedError() *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_PERMISSION_DENIED,
		Message: "Permission denied",
	}
	return e
}
