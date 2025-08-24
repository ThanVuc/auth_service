package helper

import (
	"auth_service/global"
	"auth_service/internal/grpc/models"
)

type JWTHelper interface {
	GenerateAccessToken(userId, email string, roles []string) (string, error)
	GenerateRefreshToken() string
	DecodeToken(tokenString string) (*models.JWTClaim, error)
	ValidateToken(claims models.JWTClaim) (bool, error)
	RefreshToken(refreshToken, accessToken string) (string, string, error)
	RevokeToken(accessToken, refreshToken string) error
	WriteRefreshTokenToRedis(refreshToken string) error
	WriteAccessTokenToBlacklist(jti string) error
}

func NewJWTHelper() JWTHelper {
	jwtConfig := global.Config.JWT
	return &jWTHelper{
		jwtConfig: jwtConfig,
		redis:     global.RedisDb,
	}
}
