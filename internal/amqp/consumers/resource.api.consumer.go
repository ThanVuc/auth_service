package consumers

import (
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/pkg/loggers"
	"context"
	"encoding/json"
	"os"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (cf *ConsumerFactory) InitCreateResourceApiConsumer(ctx context.Context) error {
	msgs, err := cf.consumer.RegisterConsumer(
		BrokerModeDirect,
		ExchangeNameCreateResource,
		RoutingKeyCreateResource,
		ConsumerNameCreateResource,
	)

	if err != nil {
		cf.logger.ErrorString("Failed to register consumer", zap.Error(err), zap.String("consumer", string(ConsumerNameCreateResource)))
		return err
	}

	for {
		select {
		case <-ctx.Done():
			cf.logger.InfoString("Context cancelled, stopping consumer", zap.String("consumer", string(ConsumerNameCreateResource)))
			return nil
		case msg, ok := <-msgs:
			if !ok {
				cf.logger.Warn("Message channel closed", zap.String("consumer", string(ConsumerNameCreateResource)))
				return nil
			}

			if err := cf.handleCreateResourceMessage(msg); err != nil {
				cf.logger.ErrorString("Failed to handle create resource message", zap.Error(err), zap.String("consumer", string(ConsumerNameCreateResource)))

				if ackErr := msg.Ack(false); ackErr != nil {
					cf.logger.ErrorString("Failed to acknowledge message",
						zap.Error(ackErr),
						zap.String("consumer", string(ConsumerNameCreateResource)))
				}

				continue
			}
		}
	}
}

func (cf *ConsumerFactory) handleCreateResourceMessage(msg amqp.Delivery) error {
	cf.logger.InfoString("Received message", zap.String("consumer", string(ConsumerNameCreateResource)), zap.String("message", string(msg.Body)))

	if len(msg.Body) == 0 {
		cf.logger.Warn("Received empty message body", zap.String("consumer", string(ConsumerNameCreateResource)))
		return nil
	}

	// err := CreateResourceJsonFile("resources", msg.Body, cf.logger)
	// if err != nil {
	// 	cf.logger.ErrorString("Failed to create resource JSON file", zap.Error(err), zap.String("consumer", string(ConsumerNameCreateResource)))
	// 	return err
	// }

	cf.logger.InfoString("Successfully created resource JSON file", zap.String("consumer", string(ConsumerNameCreateResource)))
	return nil
}

func CreateResourceJsonFile(service string, jsonData []byte, logger *loggers.LoggerZap) error {
	var resources []v1.ResourceItem
	err := json.Unmarshal(jsonData, &resources)
	if err != nil {
		logger.ErrorString("Failed to unmarshal JSON data", zap.Error(err))
		return err
	}

	if len(resources) == 0 {
		logger.Warn("No resources found in JSON data", zap.String("service", service))
		return nil
	}

	filePath := "./backup/" + service + "_resource.json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.ErrorString("Failed to open file for writing", zap.Error(err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Set indentation for pretty printing

	if encoder.Encode(resources); err != nil {
		logger.ErrorString("Failed to write JSON data to file", zap.Error(err))
		return err
	}

	return nil
}
