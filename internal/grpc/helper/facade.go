package helper

import (
	"auth_service/global"
	"auth_service/internal/grpc/models"
)

type JWTHelper interface {
	GenerateAccessToken(userId, email string, roles []string, jti *string) (string, error)
	GenerateRefreshToken() string
	DecodeToken(tokenString string) (*models.JWTClaim, error)
	ValidateToken(accessToken string) (*models.JWTClaim, error)
	WriteRefreshTokenToRedis(refreshToken string) error
	WriteAccessTokenToBlacklist(jti string) error
	RemoveRefreshTokenFromRedis(refreshToken string) error
}

func NewJWTHelper() JWTHelper {
	jwtConfig := global.Config.JWT
	return &jWTHelper{
		jwtConfig: jwtConfig,
		redis:     global.RedisDb,
	}
}
