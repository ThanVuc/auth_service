package controller

import (
	v1 "auth_service/internal/grpc/token.v1"
	"auth_service/internal/services"
	"context"
)

type TokenController struct {
	v1.UnimplementedTokenServiceServer
	tokenService services.ITokenService
}

func NewTokenController(
	tokenService services.ITokenService,
) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

func (tc *TokenController) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (*v1.RefreshTokenResponse, error) {
	return &v1.RefreshTokenResponse{
		Token:        "",
		RefreshToken: "",
	}, nil
}

func (tc *TokenController) RevokeToken(ctx context.Context, req *v1.RevokeTokenRequest) (*v1.RevokeTokenResponse, error) {
	return &v1.RevokeTokenResponse{
		Success: true,
	}, nil
}
