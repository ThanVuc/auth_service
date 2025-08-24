package models

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	RoleIDs []string `json:"role_ids"`
	Email   string   `json:"email"`
	jwt.RegisteredClaims
}
