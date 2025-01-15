package startup

import (
	"github.com/BevisDev/backend-template/src/main/config"
	"github.com/BevisDev/backend-template/src/main/helper"
	"sync"
)

func Run() {
	// load configuration
	startConfig()
	serverConfig := config.AppConfig.ServerConfig
	state := helper.GenUUID()
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
	defer helper.SyncAll()
	defer helper.CloseAll()
	defer helper.Close()

	// set trusted domain
	if err := r.SetTrustedProxies(serverConfig.TrustedProxies); err != nil {
		helper.LogFatal(state, "Error while setting trustedProxies: {}", err)
	}

	// run app
	if err := r.Run(serverConfig.Port); err != nil {
		helper.LogFatal(state, "Error run the server failed: {}", err)
	}
}
