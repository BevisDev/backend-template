//go:build wireinject

package di

import (
	"github.com/BevisDev/backend-template/src/main/adapter/controller"
	"github.com/BevisDev/backend-template/src/main/adapter/repositoryImpl"
	"github.com/BevisDev/backend-template/src/main/application/service/impl"
	"github.com/google/wire"
)

func NewPingDI() *controller.PingController {
	wire.Build(
		impl.NewPingServiceImpl,
		repositoryImpl.NewPingRepositoryImpl,
		controller.NewPingController,
	)
	return new(controller.PingController)
}
