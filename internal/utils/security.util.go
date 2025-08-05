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
	requestId := GetRequestIDFromOutgoingContext(ctx)
	logger := global.Logger
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic",
				requestId,
				zap.Any("error", r),
			)
		}
	}()

	return f(ctx, req)
}
