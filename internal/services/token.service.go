package services

import "auth_service/internal/repos"

type tokenService struct {
	tokenRepo repos.TokenRepo
}
