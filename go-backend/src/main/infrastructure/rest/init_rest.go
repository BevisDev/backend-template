package rest

import (
	"github.com/BevisDev/backend-template/src/main/common/utils"
	"github.com/BevisDev/backend-template/src/main/infrastructure/config"
	"github.com/BevisDev/backend-template/src/main/infrastructure/logger"
	"net/http"
	"sync"
	"time"
)

var (
	restOnce      sync.Once
	httpClient    *http.Client
	clientTimeout time.Duration
)

type Request struct {
	State  string
	URL    string
	Params map[string]any
	Header map[string]string
	Body   any
	Result any
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       any
	HasError   bool
	IsTimeout  bool
	Error      error
}

func InitRestClient(state string) *http.Client {
	cf := config.AppConfig
	if utils.IsNilOrEmpty(cf) ||
		utils.IsNilOrEmpty(cf.ServerConfig) {
		logger.Fatal(state, "Error appConfig is not initialized")
		return nil
	}
	restOnce.Do(func() {
		httpClient = &http.Client{}
		clientTimeout = time.Duration(config.AppConfig.ServerConfig.ClientTimeout) * time.Second
	})
	return httpClient
}
