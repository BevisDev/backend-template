package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
)

func Run() {
	// load configuration
	startConfig()
	serverConfig := config.AppConfig.ServerConfig
	state := utils.GenUUID()
	// logger
	startLogger(state)
	defer logger.Sync(state)

	// start app
	r := startRouter()
	//startDB(state)
	//startRedis(state)
	startRestClient()

	// set trusted domain
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
