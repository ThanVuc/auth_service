package controller

import (
	"auth_service/internal/grpc/auth"
	"auth_service/internal/services"
	"context"
)

type RoleController struct {
	auth.UnimplementedRoleServiceServer
	roleService services.RoleService
}

func NewRoleController(roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (pc *RoleController) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	response := &auth.GetRolesResponse{
		Roles: []string{},
	}
	return response, nil
}

func (pc *RoleController) CreateRole(ctx context.Context, req *auth.CreateRoleRequest) (*auth.CreateRoleResponse, error) {
	response := &auth.CreateRoleResponse{
		RoleId: "",
	}
	return response, nil
}

func (pc *RoleController) UpdateRole(ctx context.Context, req *auth.UpdateRoleRequest) (*auth.UpdateRoleResponse, error) {
	response := &auth.UpdateRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	response := &auth.DeleteRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	response := &auth.DisableOrEnableRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleRequest) (*auth.AssignRoleResponse, error) {
	response := &auth.AssignRoleResponse{
		Success: true,
	}
	return response, nil
}
