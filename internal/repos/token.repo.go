package repos

import (
	"auth_service/pkg/loggers"
	"auth_service/proto/auth"
	"context"

	"github.com/redis/go-redis/v9"
)

type tokenRepo struct {
	logger  *loggers.LoggerZap
	redisDb *redis.Client
}

func (r *tokenRepo) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) error {
	// TODO: Implement refresh token logic
	return nil
}

func (r *tokenRepo) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) error {
	// TODO: Implement revoke token logic
	return nil
}
