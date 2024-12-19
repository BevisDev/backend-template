package startup

import (
	"log"
	"time"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/BevisDev/backend-template/src/main/logger"
)

func Run() {
	// load configuration
	LoadConfig()
	serverConfig := global.AppConfig.ServerConfig

	// set time zone
	location, err := time.LoadLocation(serverConfig.Timezone)
	if err != nil {
		log.Fatalf("Error set timezone %v: %v", serverConfig.Timezone, err)
	}
	time.Local = location

	// logger
	logger.Info("LOGGER is started {}...", true)

	// router
	r := InitRouter()

	// set trusted domain
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		log.Fatalf("Error while setting trustedProxies: %v", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		log.Fatalf("Error run the server failed: %v", err)
	}
}
