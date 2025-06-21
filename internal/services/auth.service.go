package services

import (
	"auth_service/internal/repos"
)

type IAuthService interface {
}

type AuthService struct {
	authRepo repos.IAuthRepo
}

func NewAuthService(
	authRepo repos.IAuthRepo,
) IAuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}
