package helper

import (
	"auth_service/internal/grpc/models"
	"auth_service/pkg/settings"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/thanvuc/go-core-lib/cache"
)

type jWTHelper struct {
	jwtConfig settings.JWT
	redis     *cache.RedisCache
}

func (h *jWTHelper) GenerateAccessToken(userId, email string, roles []string) (string, error) {
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
			ID:        uuid.NewString(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(h.jwtConfig.Secret))
}

func (h *jWTHelper) GenerateRefreshToken() string {
	refreshToken := uuid.New().String()
	return refreshToken
}

func (h *jWTHelper) DecodeToken(tokenString string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {

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

func (h *jWTHelper) ValidateToken(claims models.JWTClaim) (bool, error) {
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return false, jwt.ErrTokenExpired
	}

	if claims.Issuer != h.jwtConfig.Issuer {
		return false, jwt.ErrTokenInvalidIssuer
	}

	key := "blacklist_token:" + claims.ID
	if exists, _ := cache.Exists(h.redis, key); exists {
		return false, jwt.ErrTokenNotValidYet
	}

	return true, nil
}

func (h *jWTHelper) RefreshToken(refreshToken, accessToken string) (string, string, error) {
	if refreshToken == "" || accessToken == "" {
		return "", "", fmt.Errorf("invalid refresh or access token")
	}

	if ok, _ := cache.Exists(h.redis, refreshToken); !ok {
		return "", "", fmt.Errorf("refresh token not found")
	}

	claim, err := h.DecodeToken(accessToken)
	if err != nil {
		return "", "", err
	}

	if claim.Issuer != h.jwtConfig.Issuer {
		return "", "", jwt.ErrTokenInvalidIssuer
	}

	key := "blacklist_token:" + claim.ID
	if exists, _ := cache.Exists(h.redis, key); exists {
		return "", "", jwt.ErrTokenNotValidYet
	}

	// can generate token when it is still valid
	newAccessToken, err := h.GenerateAccessToken(claim.Subject, claim.Email, claim.RoleIDs)
	if err != nil {
		return "", "", err
	}

	newRefreshToken := h.GenerateRefreshToken()
	if err := h.WriteRefreshTokenToRedis(newRefreshToken); err != nil {
		return "", "", err
	}

	return newRefreshToken, newAccessToken, nil
}

func (h *jWTHelper) RevokeToken(accessToken, refreshToken string) error {
	claim, err := h.DecodeToken(accessToken)
	if err != nil {
		return err
	}

	err = cache.Delete(h.redis, "refresh_token:"+refreshToken)
	if err != nil {
		return err
	}

	key := "blacklist_token:" + claim.ID
	err = cache.Set(h.redis, key, true, time.Duration(h.jwtConfig.Expiration)*time.Second)
	if err != nil {
		return err
	}

	return nil
}

func (h *jWTHelper) WriteRefreshTokenToRedis(refreshToken string) error {
	key := "refresh_token:" + refreshToken
	err := cache.Set(h.redis, key, true, time.Duration(h.jwtConfig.RefreshExp)*time.Second)
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
