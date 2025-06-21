//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
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
