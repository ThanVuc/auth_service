package global

import (
	"auth_service/pkg/loggers"
	"auth_service/pkg/settings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: This package defines global variables that are used throughout the application.
*/
var (
	Config                settings.Config
	Logger                *loggers.LoggerZap
	PostgresPool          *pgxpool.Pool
	RedisDb               *redis.Client
	RabbitMQConnection    *amqp.Connection
	RabbitMQSharedChannel *amqp.Channel
)
