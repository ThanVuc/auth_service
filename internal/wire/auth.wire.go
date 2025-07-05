//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/internal/services"

	"github.com/google/wire"
)

func InjectAuthWire() *controller.AuthController {
	wire.Build(
		repos.NewAuthRepo,
		services.NewAuthService,
		controller.NewAuthController,
	)

	return nil
}

func InjectPermissionWire() *controller.PermissionController {
	wire.Build(
		repos.NewPermissionRepo,
		mapper.NewPermissionMapper,
		services.NewPermissionService,
		controller.NewPermissionController,
	)

	return nil
}

func InjectRoleWire() *controller.RoleController {
	wire.Build(
		repos.NewRoleRepo,
		services.NewRoleService,
		controller.NewRoleController,
	)

	return nil
}

func InjectTokenWire() *controller.TokenController {
	wire.Build(
		repos.NewTokenRepo,
		services.NewTokenService,
		controller.NewTokenController,
	)

	return nil
}
