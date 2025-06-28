package controller

import (
	"auth_service/internal/grpc/auth"
	"auth_service/internal/services"
	"auth_service/internal/utils"
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
	response := &auth.GetPermissionsResponse{
		Permissions: []string{},
	}
	return response, nil
}

func (pc *PermissionController) CreatePermission(ctx context.Context, req *auth.CreatePermissionRequest) (*auth.CreatePermissionResponse, error) {
	response := &auth.CreatePermissionResponse{
		PermissionId: "",
	}
	return response, nil
}

func (pc *PermissionController) UpdatePermission(ctx context.Context, req *auth.UpdatePermissionRequest) (*auth.UpdatePermissionResponse, error) {
	response := &auth.UpdatePermissionResponse{
		Success: true,
	}
	return response, nil
}

func (pc *PermissionController) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	response := &auth.DeletePermissionResponse{
		Success: true,
	}
	return response, nil
}

func (pc *PermissionController) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	response := &auth.AssignPermissionResponse{
		Success: true,
	}
	return response, nil
}

func (pc *PermissionController) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	return utils.WithSafePanic(ctx, req, pc.permissionService.GetResources)
}
