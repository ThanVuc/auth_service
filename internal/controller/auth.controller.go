package controller

import (
	"auth_service/global"
	"auth_service/proto/auth"
	"auth_service/proto/common"

	"auth_service/internal/services"
	"auth_service/internal/utils"
	"auth_service/pkg/loggers"
	"context"
)

type AuthController struct {
	auth.UnimplementedAuthServiceServer
	authService services.AuthService
	logger      *loggers.LoggerZap
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
		logger:      global.Logger,
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
