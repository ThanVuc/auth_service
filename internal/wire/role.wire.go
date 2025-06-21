//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
	"auth_service/internal/repos"
	"auth_service/internal/services"

	"github.com/google/wire"
)

func InjectRoleWire() *controller.RoleController {
	wire.Build(
		repos.NewRoleRepo,
		services.NewRoleService,
		controller.NewRoleController,
	)

	return nil
}
