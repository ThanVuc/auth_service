package utils

import "auth_service/proto/common"

var ErrorMessage = map[string]string{
	"DatabaseError": "Database operation failed",
	"NotFoundError": "Resource not found",
	"RuntimeError":  "An unexpected error occurred",
}

func DatabaseError(detail string) *common.Error {
	err := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_DATABASE_ERROR,
		Message: "Database operation failed: " + detail,
	}
	return err
}

func NotFoundError(detail string) *common.Error {
	err := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_NOT_FOUND,
		Message: "Resource not found: " + detail,
	}
	return err
}

func RuntimeError(detail string) *common.Error {
	err := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_RUN_TIME_ERROR,
		Message: "An unexpected error occurred: " + detail,
	}
	return err
}

func UnorthorizedError(detail string) *common.Error {
	err := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_UNAUTHORIZED,
		Message: "Unauthorized access: " + detail,
	}
	return err
}

func PermissionDeniedError(detail string) *common.Error {
	err := &common.Error{
		Code:    common.ErrorCode_ERROR_CODE_PERMISSION_DENIED,
		Message: "Permission denied: " + detail,
	}
	return err
}
