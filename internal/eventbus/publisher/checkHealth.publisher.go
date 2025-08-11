package publisher

import (
	"auth_service/global"
	"context"

	"github.com/google/uuid"
	"github.com/thanvuc/go-core-lib/eventbus"
)

func PushCheckHealthMessage(ctx context.Context) {
	requestID := uuid.NewString()
	maxRetries := 3
	retryDelay := 3
	dlqExchange := eventbus.DLQCheckHealthExchange
	dlqRoutingKey := "check_health_dlq"

	publisher := eventbus.NewPublisher(
		global.EventBusConnector,
		eventbus.CheckHealthExchange,
		[]string{"check_health"},
		&maxRetries,
		&retryDelay,
		&dlqExchange,
		&dlqRoutingKey,
	)

	err := publisher.SafetyPublish(
		ctx,
		requestID,
		[]byte("check_health"),
		nil,
	)

	if err != nil {
		global.Logger.Error("Failed to publish check health message", requestID)
		return
	}

	global.Logger.Info("Check health message published successfully", requestID)
}
