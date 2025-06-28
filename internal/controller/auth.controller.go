package controller

import (
	"auth_service/global"
	"auth_service/internal/grpc/auth"
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
	// Call the authService to handle login logic
	response := &auth.LoginResponse{}
	if req.Username == "sinhnguyen" && req.Password == "123456" {
		response.AccessToken = "access_token_example"
		response.RefreshToken = "refresh_token_example"
	}

	return response, nil
}

func (ac *AuthController) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	response := &auth.RegisterResponse{
		IsSentMail: true,
	}
	return response, nil
}

func (ac *AuthController) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error) {
	response := &auth.ConfirmEmailResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return response, nil
}

func (ac *AuthController) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	response := &auth.LogoutResponse{
		Success: true,
	}
	return response, nil
}

// forgot password
func (ac *AuthController) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
	response := &auth.ForgotPasswordResponse{
		UserId: "12345",
	}
	return response, nil
}

func (ac *AuthController) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error) {
	response := &auth.ConfirmForgotPasswordResponse{
		Success:      true,
		AccessToken:  "",
		RefreshToken: "",
		PasswordTmp:  "",
	}
	return response, nil
}

func (ac *AuthController) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	response := &auth.ResetPasswordResponse{
		Success: true,
	}
	return response, nil
}

func (ac *AuthController) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	return utils.WithSafePanic(ctx, req, ac.authService.SaveRouteResource)
}
