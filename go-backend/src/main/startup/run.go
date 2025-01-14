package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper/db"
	"github.com/BevisDev/backend-template/src/main/helper/logger"
	"github.com/BevisDev/backend-template/src/main/helper/redis"
	"github.com/BevisDev/backend-template/src/main/helper/utils"
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
	defer db.CloseAll()
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
