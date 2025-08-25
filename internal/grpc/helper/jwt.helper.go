package helper

import (
	"auth_service/internal/grpc/models"
	"auth_service/pkg/settings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/thanvuc/go-core-lib/cache"
)

type jWTHelper struct {
	jwtConfig settings.JWT
	redis     *cache.RedisCache
}

func (h *jWTHelper) GenerateAccessToken(userId, email string, roles []string, jti *string) (string, error) {
	claims := models.JWTClaim{
		Email:   email,
		RoleIDs: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    h.jwtConfig.Issuer,
			Audience:  h.jwtConfig.Audience,
			Subject:   userId,
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(h.jwtConfig.Expiration) * time.Second)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			NotBefore: &jwt.NumericDate{Time: time.Now()},
		},
	}

	if jti != nil {
		claims.ID = *jti
	} else {
		claims.ID = uuid.New().String()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.jwtConfig.Secret))
}

func (h *jWTHelper) GenerateRefreshToken() string {
	refreshToken := uuid.New().String()
	return refreshToken
}

func (h *jWTHelper) DecodeToken(accessToken string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(h.jwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func (h *jWTHelper) ValidateToken(accessToken string) (*models.JWTClaim, error) {
	claims, err := h.DecodeToken(accessToken)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	if claims.Issuer != h.jwtConfig.Issuer {
		return nil, jwt.ErrTokenInvalidIssuer
	}

	key := "blacklist_token:" + claims.ID
	if exists, _ := cache.Exists(h.redis, key); exists {
		return nil, jwt.ErrTokenNotValidYet
	}

	return claims, nil
}

func (h *jWTHelper) WriteRefreshTokenToRedis(refreshToken string) error {
	key := "refresh_token:" + refreshToken
	err := cache.Set(h.redis, key, true, time.Duration(h.jwtConfig.RefreshExp)*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (h *jWTHelper) RemoveRefreshTokenFromRedis(refreshToken string) error {
	key := "refresh_token:" + refreshToken
	err := cache.Delete(h.redis, key)
	if err != nil {
		return err
	}

	return nil
}

func (h *jWTHelper) WriteAccessTokenToBlacklist(jti string) error {
	key := "blacklist_token:" + jti
	err := cache.Set(h.redis, key, true, time.Duration(h.jwtConfig.Expiration)*time.Second)
	if err != nil {
		return err
	}
	return nil
}
