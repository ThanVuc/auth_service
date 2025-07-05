package controller

import (
	"auth_service/global"
	"auth_service/proto/auth"

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

func (ac *AuthController) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.Login)
}

func (ac *AuthController) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.Register)
}

func (ac *AuthController) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.ConfirmEmail)
}

func (ac *AuthController) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.Logout)
}

// forgot password
func (ac *AuthController) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.ForgotPassword)
}

func (ac *AuthController) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.ConfirmForgotPassword)
}

func (ac *AuthController) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.ResetPassword)
}

func (ac *AuthController) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.SaveRouteResource)
}
