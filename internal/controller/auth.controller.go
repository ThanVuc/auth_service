package controller

import (
	v1 "auth_service/internal/grpc/auth.v1"
	"auth_service/internal/services"
	"context"
)

type AuthController struct {
	v1.UnimplementedAuthServiceServer
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// Call the authService to handle login logic

	response := &v1.LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
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
		Success: true,
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
