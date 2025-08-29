//go:build wireinject

package wire

import (
	"auth_service/internal/grpc/controller"
	"auth_service/internal/grpc/helper"
	"auth_service/internal/grpc/mapper"
	"auth_service/internal/grpc/repos"
	"auth_service/internal/grpc/services"

	"github.com/google/wire"
)

func InjectAuthController() *controller.AuthController {
	wire.Build(
		repos.NewAuthRepo,
		helper.NewJWTHelper,
		mapper.NewAuthMapper,
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

func InjectUserController() *controller.UserController {
	wire.Build(
		repos.NewUserRepo,
		mapper.NewUserMapper,
		services.NewUserService,
		controller.NewUserController,
	)

	return nil
}
