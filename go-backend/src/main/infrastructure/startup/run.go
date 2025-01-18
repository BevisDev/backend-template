package startup

import (
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/database"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"github.com/BevisDev/backend-template/src/main/infrastructure/redis"
	"sync"
)

func Run() {
	// load configuration
	startConfig()
	serverConfig := config.AppConfig.ServerConfig
	state := utils.GenUUID()
	// logger
	startLogger(state)
	// start app
	r := startRouter()
	var wg sync.WaitGroup
	// rest client
	wg.Add(1)
	go func() {
		defer wg.Done()
		startRestClient(state)
	}()
	//// db
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	startDB(state)
	//}()
	//// redis
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	startRedis(state)
	//}()

	wg.Wait()
	defer logger.SyncAll()
	defer database.CloseAll()
	defer redis.Close()

	// set trusted domain
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		logger.Fatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		logger.Fatal(state, "Error run the server failed: {}", err)
	}
}
