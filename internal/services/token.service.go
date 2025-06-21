package services

import "auth_service/internal/repos"

type ITokenService interface{}

type TokenService struct {
	tokenRepo repos.ITokenRepo
}

func NewTokenService(tokenRepo repos.ITokenRepo) ITokenService {
	return &TokenService{
		tokenRepo: tokenRepo,
	}
}
