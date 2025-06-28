package utils

import (
	"auth_service/global"
	"context"

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
			logger.ErrorString("Recovered from panic", zap.Any("error", r), zap.Any("request", req))
		}
	}()
	return f(ctx, req)
}
