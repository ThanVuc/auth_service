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

	resp, err := f(ctx, req)
	if err != nil {
		logger.Error("Error occurred in WithSafePanic",
			requestId,
			zap.Error(err),
		)
		return resp, err
	}

	return resp, err
}
