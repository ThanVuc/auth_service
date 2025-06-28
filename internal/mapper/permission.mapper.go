package mapper

import (
	"auth_service/internal/database"
	"auth_service/internal/grpc/auth"
)

type permissionMapper struct{}

func (p *permissionMapper) ConvertDbResourcesRowToGrpcResources(resources []database.GetResourcesRow) []*auth.Resource {
	result := make([]*auth.Resource, 0)
	for _, resource := range resources {
		result = append(result, &auth.Resource{
			Id:   resource.ResourceID,
			Name: resource.Name,
		})
	}

	return result
}
