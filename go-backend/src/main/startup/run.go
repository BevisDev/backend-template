package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"time"
)

func Run() {
	// load configuration
	LoadConfig()
	serverConfig := config.AppConfig.ServerConfig

	// init logger
	state := utils.GenUUID()
	logger.Info(state, "LOGGER is started {}...", true)

	// Defer Sync to ensure logs are flushed before exiting
	defer logger.Sync(state)
	// recover global when occur exception
	defer func() {
		if r := recover(); r != nil {
			logger.Error(state, "Recovered in run.go:20 with {}", r)
		}
	}()

	// set time zone
	location, err := time.LoadLocation(serverConfig.Timezone)
	if err != nil {
		logger.Panic("Error set timezone {}: {}", serverConfig.Timezone, err)
	}
	time.Local = location

	// router
	r := InitRouter()

	// set trusted domain
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		logger.Panic(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		logger.Panic(state, "Error run the server failed: {}", err)
	}
}
