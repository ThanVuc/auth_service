package initialize

import (
	"auth_service/internal/amqp/consumers"
	"auth_service/pkg/loggers"
	"context"

	"go.uber.org/zap"
)

func InitAllConsumers(ctx context.Context, logger *loggers.LoggerZap) {
	consumerFactory := consumers.NewConsumerFactory()

	go func() {
		if err := consumerFactory.InitCreateResourceApiConsumer(ctx); err != nil {
			logger.ErrorString("Failed to initialize CreateResourceApiConsumer", zap.Error(err))
			return
		}
	}()

}
