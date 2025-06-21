package initialize

import (
	"auth_service/global"
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitRabbitMQ() error {
	var logger = global.Logger
	var config = global.Config.RabbitMQ

	var err error
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.ErrorString("mq.GetConnection: failed to connect to RabbitMQ", zap.String("error", err.Error()))
		return err
	}

	// Ensure global.RabbitMQ is initialized

	// Set the connection to the global variable
	global.RabbitMQConnection = conn
	global.RabbitMQSharedChannel, err = global.RabbitMQConnection.Channel()

	if err != nil {
		logger.ErrorString("mq.GetConnection: failed to open a channel", zap.String("error", err.Error()))
		return err
	}

	logger.InfoString("mq.GetConnection: RabbitMQ connection established successfully")
	return nil
}

func CloseConnection() error {
	var logger = global.Logger
	if global.RabbitMQConnection != nil {
		err := global.RabbitMQConnection.Close()
		if err != nil {
			logger.ErrorString("mq.CloseConnection: failed to close RabbitMQ connection", zap.String("error", err.Error()))
			return err
		}
		logger.InfoString("mq.CloseConnection: RabbitMQ connection closed successfully")
		return nil
	}

	logger.InfoString("mq.CloseConnection: no RabbitMQ connection to close")
	return nil
}
