package repos

import (
	"auth_service/pkg/loggers"

	"github.com/redis/go-redis/v9"
)

type TokenRepo struct {
	logger  *loggers.LoggerZap
	redisDb *redis.Client
}
