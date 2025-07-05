package services

import (
	"auth_service/internal/repos"
	"auth_service/proto/auth"
	"context"
)

type roleService struct {
	repo repos.RoleRepo
}

func (rs *roleService) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	// TODO: Implement get roles logic
	return &auth.GetRolesResponse{}, nil
}

func (rs *roleService) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error) {
	// TODO: Implement upsert role logic
	return &auth.UpsertRoleResponse{}, nil
}

func (rs *roleService) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	// TODO: Implement delete role logic
	return &auth.DeleteRoleResponse{}, nil
}

func (rs *roleService) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	// TODO: Implement disable/enable role logic
	return &auth.DisableOrEnableRoleResponse{}, nil
}

func (rs *roleService) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleRequest) (*auth.AssignRoleResponse, error) {
	// TODO: Implement assign role to user logic
	return &auth.AssignRoleResponse{}, nil
}
