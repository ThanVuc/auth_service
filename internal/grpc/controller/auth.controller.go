package controller

import (
	"auth_service/internal/grpc/services"
	"auth_service/internal/grpc/utils"
	"auth_service/proto/auth"
	"auth_service/proto/common"
	"context"
)

type AuthController struct {
	auth.UnimplementedAuthServiceServer
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) LoginWithGoogle(ctx context.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.LoginWithGoogle)
}

func (ac *AuthController) Logout(ctx context.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.Logout)
}

func (ac *AuthController) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.SaveRouteResource)
}
