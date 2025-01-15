// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/BevisDev/backend-template/src/main/controller"
	"github.com/BevisDev/backend-template/src/main/repository"
	"github.com/BevisDev/backend-template/src/main/service/impl"
)

// Injectors from ping_di.go:

func NewPingDI() *controller.PingController {
	iPingRepository := repository.NewPingRepository()
	iPingService := impl.NewPingServiceImpl(iPingRepository)
	pingController := controller.NewPingController(iPingService)
	return pingController
}
