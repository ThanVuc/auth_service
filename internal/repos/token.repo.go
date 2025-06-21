package repos

import (
	"auth_service/global"

	"github.com/redis/go-redis/v9"
)

type ITokenRepo interface {
}

type TokenRepo struct {
	redisDb *redis.Client
}

func NewTokenRepo() ITokenRepo {
	return &TokenRepo{
		redisDb: global.RedisDb,
	}
}
