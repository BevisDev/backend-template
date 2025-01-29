//go:build wireinject

package di

import (
	"github.com/BevisDev/backend-template/src/main/adapter/controller"
	"github.com/BevisDev/backend-template/src/main/application/service/impl"
	"github.com/google/wire"
)

func NewAuthDI() *controller.AuthController {
	wire.Build(
		impl.NewAuthServiceImpl,
		controller.NewAuthController,
	)
	return new(controller.AuthController)
}
