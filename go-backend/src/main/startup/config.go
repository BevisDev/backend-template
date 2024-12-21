package startup

import (
	"log"
	"os"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	profile := os.Getenv("GO_PROFILE")
	if profile == "" {
		profile = "dev" // set default
	}

	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigName(profile)
	v.SetConfigType("yaml")

	// read config
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error can not read configuration %v", err)
	}

	if err := v.Unmarshal(&global.AppConfig); err != nil {
		log.Fatalf("Error decode config into struct, %v", err)
	}

	serverConfig := global.AppConfig.ServerConfig
	log.Println("================================")
	log.Printf("Load configuration profile %s successful", serverConfig.Profile)
	log.Printf("Welcome to %s version %s ", serverConfig.Name, serverConfig.Version)
	log.Println("================================")
}
