//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
	"auth_service/internal/repos"
	"auth_service/internal/services"

	"github.com/google/wire"
)

func InjectPermissionWire() *controller.PermissionController {
	wire.Build(
		repos.NewPermissionRepo,
		services.NewPermissionService,
		controller.NewPermissionController,
	)

	return nil
}
