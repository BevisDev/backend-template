package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
	_ "github.com/BevisDev/backend-template/src/resources/swagger"
)

// @title           API Specification
// @version         1.0
// @description     There are APIs in project
// @termsOfService  https://github.com/BevisDev

// @contact.name   Truong Thanh Binh
// @contact.url    https://github.com/BevisDev
// @contact.email  dev.binhtt@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8089
// @BasePath  /api

// @securityDefinitions.apiKey AccessTokenAuth
// @in 								header
// @name	 						AccessToken
// @description						Description for what is this security definition being used

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
