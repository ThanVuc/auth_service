package controller

import (
	"auth_service/internal/services"
	"auth_service/internal/utils"
	"auth_service/proto/auth"
	"context"
)

type PermissionController struct {
	auth.UnimplementedPermissionServiceServer
	permissionService services.PermissionService
}

func NewPermissionController(permissionService services.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

func (pc *PermissionController) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.GetPermissions)
}

func (pc *PermissionController) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.UpsertPermission)
}

func (pc *PermissionController) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.DeletePermission)
}

func (pc *PermissionController) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.AssignPermissionToRole)
}

func (pc *PermissionController) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.GetResources)
}

func (pc *PermissionController) GetActions(ctx context.Context, req *auth.GetActionsRequest) (*auth.GetActionsResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.GetActions)
}
