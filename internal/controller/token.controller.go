package controller

import (
	"auth_service/internal/grpc/auth"
	"auth_service/internal/services"
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
	return &auth.RefreshTokenResponse{
		Token:        "",
		RefreshToken: "",
	}, nil
}

func (tc *TokenController) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	return &auth.RevokeTokenResponse{
		Success: true,
	}, nil
}
