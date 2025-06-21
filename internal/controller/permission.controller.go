package controller

import (
	v1 "auth_service/internal/grpc/permission.v1"
	"auth_service/internal/services"
	"context"
)

type PermissionController struct {
	v1.UnimplementedPermissionServiceServer
	permissionService services.IPermissionService
}

func NewPermissionController(permissionService services.IPermissionService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

func (pc *PermissionController) GetPermissions(ctx context.Context, req *v1.GetPermissionsRequest) (*v1.GetPermissionsResponse, error) {
	response := &v1.GetPermissionsResponse{
		Permissions: []string{},
	}
	return response, nil
}

func (pc *PermissionController) CreatePermission(ctx context.Context, req *v1.CreatePermissionRequest) (*v1.CreatePermissionResponse, error) {
	response := &v1.CreatePermissionResponse{
		PermissionId: "",
	}
	return response, nil
}

func (pc *PermissionController) UpdatePermission(ctx context.Context, req *v1.UpdatePermissionRequest) (*v1.UpdatePermissionResponse, error) {
	response := &v1.UpdatePermissionResponse{
		Success: true,
	}
	return response, nil
}

func (pc *PermissionController) DeletePermission(ctx context.Context, req *v1.DeletePermissionRequest) (*v1.DeletePermissionResponse, error) {
	response := &v1.DeletePermissionResponse{
		Success: true,
	}
	return response, nil
}

func (pc *PermissionController) AssignPermissionToRole(ctx context.Context, req *v1.AssignPermissionRequest) (*v1.AssignPermissionResponse, error) {
	response := &v1.AssignPermissionResponse{
		Success: true,
	}
	return response, nil
}
