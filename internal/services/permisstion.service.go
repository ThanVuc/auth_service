package services

import "auth_service/internal/repos"

type IPermissionService interface{}

type PermissionService struct {
	permissionRepo repos.IPermissionRepo
}

func NewPermissionService(permissionRepo repos.IPermissionRepo) IPermissionService {
	return &PermissionService{
		permissionRepo: permissionRepo,
	}
}
