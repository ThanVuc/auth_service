package utils

import (
	"auth_service/global"
	"context"
	"fmt"
	"runtime/debug"

	"go.uber.org/zap"
)

func WithSafePanic[TReq any, TResp any](
	ctx context.Context,
	req TReq,
	f func(context.Context, TReq) (TResp, error),
) (TResp, error) {
	defer func() {
		if r := recover(); r != nil {
			logger := global.Logger
			stack := string(debug.Stack())
			logger.ErrorString("Recovered from panic", zap.Any("error", r),
				zap.Any("request", req),
				zap.String("stacktrace", stack))

			fmt.Println("stacktrace: \n" + stack)
		}
	}()
	return f(ctx, req)
}
