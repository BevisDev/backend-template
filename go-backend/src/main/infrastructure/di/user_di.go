package di

import (
	"github.com/BevisDev/backend-template/src/main/adapter/controller"
	"github.com/BevisDev/backend-template/src/main/application/service/impl"
	"github.com/google/wire"
)

func NewUserDI() *controller.UserController {
	wire.Build(
		impl.NewUserServiceImpl,
		controller.NewUserController,
	)
	return new(controller.UserController)
}
