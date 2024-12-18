package startup

import (
	"log"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/BevisDev/backend-template/src/main/logger"
)

func Run() {
	// load configuration
	LoadConfig()

	// logger
	logger.NewLogger()

	// router
	r := InitRouter()

	// set trusted domain
	serverConfig := global.AppConfig.ServerConfig
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		log.Fatalf("Error while setting trustedProxies: %v", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		log.Fatalf("Error run the server failed: %v", err)
	}
}
