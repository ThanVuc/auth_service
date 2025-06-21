package services

import "auth_service/internal/repos"

type IRoleService interface {
}

type RoleService struct {
	repo repos.IRoleRepo
}

func NewRoleService(roleRepo repos.IRoleRepo) IRoleService {
	return &RoleService{
		repo: roleRepo,
	}
}
