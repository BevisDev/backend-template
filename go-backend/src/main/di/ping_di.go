//go:build wireinject

package di

import (
	"github.com/BevisDev/backend-template/src/main/controller"
	"github.com/BevisDev/backend-template/src/main/repository"
	"github.com/BevisDev/backend-template/src/main/service/impl"
	"github.com/google/wire"
)

func NewPingDI() *controller.PingController {
	wire.Build(
		impl.NewPingServiceImpl,
		repository.NewPingRepository,
		controller.NewPingController,
	)
	return new(controller.PingController)
}
