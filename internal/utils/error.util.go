package utils

import (
	"auth_service/global"
	"auth_service/proto/common"
	"runtime/debug"

	"go.uber.org/zap"
)

var ErrorMessage = map[string]string{
	"DatabaseError": "Database operation failed",
	"NotFoundError": "Resource not found",
	"RuntimeError":  "An unexpected error occurred",
}

func DatabaseError(err error) *common.Error {
	message := "Database operation failed"
	if err != nil {
		message += ": " + err.Error()
	}
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_DATABASE_ERROR,
		Message: message,
	}
	writeTrace(err)
	return e
}

func NotFoundError(err error) *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_NOT_FOUND,
		Message: "Not found error",
	}
	writeTrace(err)
	return e
}

func RuntimeError(err error) *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_RUN_TIME_ERROR,
		Message: "An unexpected error occurred: runtime error",
	}
	writeTrace(err)
	return e
}

func UnauthorizedError(err error) *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_UNAUTHORIZED,
		Message: "Unauthorized access",
	}
	writeTrace(err)
	return e
}

func PermissionDeniedError(err error) *common.Error {
	e := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_PERMISSION_DENIED,
		Message: "Permission denied",
	}
	writeTrace(err)
	return e
}

func writeTrace(err error) {
	if err != nil {
		global.Logger.ErrorString("Error occurred", zap.Error(err))
		debug.PrintStack()
	}
}
