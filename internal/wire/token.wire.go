//go:build wireinject

package wire

import (
	"auth_service/internal/controller"
	"auth_service/internal/repos"
	"auth_service/internal/services"

	"github.com/google/wire"
)

func InjectTokenWire() *controller.TokenController {
	wire.Build(
		repos.NewTokenRepo,
		services.NewTokenService,
		controller.NewTokenController,
	)

	return nil
}
