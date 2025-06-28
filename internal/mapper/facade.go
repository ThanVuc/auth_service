package mapper

import (
	"auth_service/internal/database"
	"auth_service/internal/grpc/auth"
)

type (
	PermissionMapper interface {
		ConvertDbResourcesRowToGrpcResources(resources []database.GetResourcesRow) []*auth.Resource
	}
)

func NewPermissionMapper() PermissionMapper {
	return &permissionMapper{}
}
