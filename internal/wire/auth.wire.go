//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
	"auth_service/internal/mapper"
	"auth_service/internal/repos"
	"auth_service/internal/services"

	"github.com/google/wire"
)

func InjectAuthController() *controller.AuthController {
	wire.Build(
		repos.NewAuthRepo,
		services.NewAuthService,
		controller.NewAuthController,
	)

	return nil
}

func InjectPermissionController() *controller.PermissionController {
	wire.Build(
		repos.NewPermissionRepo,
		mapper.NewPermissionMapper,
		services.NewPermissionService,
		controller.NewPermissionController,
	)

	return nil
}

func InjectRoleController() *controller.RoleController {
	wire.Build(
		repos.NewRoleRepo,
		mapper.NewRoleMapper,
		services.NewRoleService,
		controller.NewRoleController,
	)

	return nil
}

func InjectTokenController() *controller.TokenController {
	wire.Build(
		repos.NewTokenRepo,
		services.NewTokenService,
		controller.NewTokenController,
	)

	return nil
}
