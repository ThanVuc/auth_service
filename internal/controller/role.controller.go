package controller

import (
	v1 "auth_service/internal/grpc/role.v1"
	"auth_service/internal/services"
	"context"
)

type RoleController struct {
	v1.UnimplementedRoleServiceServer
	roleService services.IRoleService
}

func NewRoleController(roleService services.IRoleService) *RoleController {
	return &RoleController{
		roleService: roleService,
	}
}

func (pc *RoleController) GetRoles(ctx context.Context, req *v1.GetRolesRequest) (*v1.GetRolesResponse, error) {
	response := &v1.GetRolesResponse{
		Roles: []string{},
	}
	return response, nil
}

func (pc *RoleController) CreateRole(ctx context.Context, req *v1.CreateRoleRequest) (*v1.CreateRoleResponse, error) {
	response := &v1.CreateRoleResponse{
		RoleId: "",
	}
	return response, nil
}

func (pc *RoleController) UpdateRole(ctx context.Context, req *v1.UpdateRoleRequest) (*v1.UpdateRoleResponse, error) {
	response := &v1.UpdateRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) DeleteRole(ctx context.Context, req *v1.DeleteRoleRequest) (*v1.DeleteRoleResponse, error) {
	response := &v1.DeleteRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) DisableOrEnableRole(ctx context.Context, req *v1.DisableOrEnableRoleRequest) (*v1.DisableOrEnableRoleResponse, error) {
	response := &v1.DisableOrEnableRoleResponse{
		Success: true,
	}
	return response, nil
}

func (pc *RoleController) AssignRoleToUser(ctx context.Context, req *v1.AssignRoleRequest) (*v1.AssignRoleResponse, error) {
	response := &v1.AssignRoleResponse{
		Success: true,
	}
	return response, nil
}
