package mapper

import (
	"auth_service/internal/grpc/database"
	"auth_service/proto/auth"
)

type authMapper struct{}

func (am *authMapper) ConvertFromUserAuthRowToProto(from []database.GetUserAuthInfoRow) []*auth.PermissionAuthItem {
	permissionMap := make(map[string]*auth.PermissionAuthItem)
	for _, item := range from {
		permId := item.PermID.String()
		if _, exists := permissionMap[permId]; !exists {
			permissionMap[permId] = &auth.PermissionAuthItem{
				Permission: item.PermName,
				Resource:   item.ResourceName,
				Actions:    make([]string, 0),
			}
		}

		permissionMap[permId].Actions = append(permissionMap[permId].Actions, item.ActionName)
	}

	result := make([]*auth.PermissionAuthItem, 0, len(permissionMap))
	for _, perm := range permissionMap {
		result = append(result, perm)
	}

	return result
}
