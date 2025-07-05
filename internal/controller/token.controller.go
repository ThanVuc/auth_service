package controller

import (
	"auth_service/internal/services"
	"auth_service/internal/utils"
	"auth_service/proto/auth"
	"context"
)

type TokenController struct {
	auth.UnimplementedTokenServiceServer
	tokenService services.TokenService
}

func NewTokenController(
	tokenService services.TokenService,
) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

func (tc *TokenController) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	return utils.WithSafePanic(ctx, req, tc.tokenService.RefreshToken)
}

func (tc *TokenController) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	return utils.WithSafePanic(ctx, req, tc.tokenService.RevokeToken)
}
