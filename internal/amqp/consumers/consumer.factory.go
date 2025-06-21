package consumers

import (
	"auth_service/global"
	"auth_service/pkg/loggers"
	"context"
)

type IConsumerFactory interface {
	InitCreateResourceApiConsumer(ctx context.Context) error
}

type ConsumerFactory struct {
	consumer IConsumer
	logger   *loggers.LoggerZap
}

func NewConsumerFactory() IConsumerFactory {
	return &ConsumerFactory{
		consumer: NewConsumer(),
		logger:   global.Logger,
	}
}
