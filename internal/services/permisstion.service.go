package services

import (
	"auth_service/internal/grpc/auth"
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"context"
)

type permissionService struct {
	permissionRepo   repos.PermissionRepo
	permissionMapper mapper.PermissionMapper
}

func (ps *permissionService) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	dbResource, err := ps.permissionRepo.GetResources(ctx)
	if err != nil {
		return nil, err
	}

	resourceItem := ps.permissionMapper.ConvertDbResourcesRowToGrpcResources(dbResource)
	resp := &auth.GetResourcesResponse{
		Resources: resourceItem,
	}

	return resp, nil
}
