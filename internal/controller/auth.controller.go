package controller

import (
	"auth_service/global"
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/internal/services"
	"auth_service/pkg/loggers"
	"context"
	"fmt"

	"go.uber.org/zap"
)

type AuthController struct {
	v1.UnimplementedAuthServiceServer
	authService services.IAuthService
	logger      *loggers.LoggerZap
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
		logger:      global.Logger,
	}
}

func (ac *AuthController) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// Call the authService to handle login logic
	response := &v1.LoginResponse{}
	if req.Username == "sinhnguyen" && req.Password == "123456" {
		response.AccessToken = "access_token_example"
		response.RefreshToken = "refresh_token_example"
	}

	return response, nil
}

func (ac *AuthController) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	response := &v1.RegisterResponse{
		IsSentMail: true,
	}
	return response, nil
}

func (ac *AuthController) ConfirmEmail(ctx context.Context, req *v1.ConfirmEmailRequest) (*v1.ConfirmEmailResponse, error) {
	response := &v1.ConfirmEmailResponse{
		AccessToken:  "",
		RefreshToken: "",
	}
	return response, nil
}

func (ac *AuthController) Logout(ctx context.Context, req *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	response := &v1.LogoutResponse{
		Success: true,
	}
	return response, nil
}

// forgot password
func (ac *AuthController) ForgotPassword(ctx context.Context, req *v1.ForgotPasswordRequest) (*v1.ForgotPasswordResponse, error) {
	response := &v1.ForgotPasswordResponse{
		UserId: "12345",
	}
	return response, nil
}

func (ac *AuthController) ConfirmForgotPassword(ctx context.Context, req *v1.ConfirmForgotPasswordRequest) (*v1.ConfirmForgotPasswordResponse, error) {
	response := &v1.ConfirmForgotPasswordResponse{
		Success:      true,
		AccessToken:  "",
		RefreshToken: "",
		PasswordTmp:  "",
	}
	return response, nil
}

func (ac *AuthController) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.ResetPasswordResponse, error) {
	response := &v1.ResetPasswordResponse{
		Success: true,
	}
	return response, nil
}

func (ac *AuthController) SaveRouteResource(ctx context.Context, req *v1.SaveRouteResourceRequest) (*v1.SaveRouteResourceResponse, error) {
	items := req.Items
	if len(items) == 0 {
		ac.logger.ErrorString("SaveRouteResource: no items provided", zap.String("request", fmt.Sprintf("%+v", req)))
		return nil, fmt.Errorf("no items provided")
	}

	result := ac.authService.SaveRouteResource(items)

	if !result {
		ac.logger.ErrorString("SaveRouteResource: failed to save route resources", zap.String("request", fmt.Sprintf("%+v", req)))
		return nil, fmt.Errorf("failed to save route resources")
	}

	ac.logger.InfoString("SaveRouteResource: route resources saved successfully")

	response := &v1.SaveRouteResourceResponse{
		Success: true,
		Message: "Route resource saved successfully",
	}
	return response, nil
}
