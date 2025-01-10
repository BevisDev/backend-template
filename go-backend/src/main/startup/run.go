package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	"time"
)

func Run() {
	// load configuration
	loadConfig()
	serverConfig := config.AppConfig.ServerConfig
	state := utils.GenUUID()

	// logger
	startLogger(state)

	defer logger.Sync(state)

	// set time zone
	location, err := time.LoadLocation(serverConfig.Timezone)
	if err != nil {
		logger.Fatal(state, "Error set timezone {}: {}", serverConfig.Timezone, err)
	}
	time.Local = location

	// router
	r := startRouter()
	// db
	startDB(state)
	// redis
	startRedis(state)
	// rest
	startRestClient()

	// set trusted domain
	if err = r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err = r.Run(serverConfig.Port); err != nil {
		logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
