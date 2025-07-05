package services

import (
	"auth_service/internal/repos"
	"auth_service/proto/auth"
	"context"
)

type tokenService struct {
	tokenRepo repos.TokenRepo
}

func (ts *tokenService) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	// TODO: Implement refresh token logic
	return &auth.RefreshTokenResponse{}, nil
}

func (ts *tokenService) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	// TODO: Implement revoke token logic
	return &auth.RevokeTokenResponse{}, nil
}
