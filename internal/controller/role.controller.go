package controller

import (
	"auth_service/internal/services"
	"auth_service/internal/utils"
	"auth_service/proto/auth"
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
	return utils.WithSafePanic(ctx, req, pc.roleService.GetRoles)
}

func (pc *RoleController) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.roleService.UpsertRole)
}

func (pc *RoleController) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.roleService.DeleteRole)
}

func (pc *RoleController) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.roleService.DisableOrEnableRole)
}
