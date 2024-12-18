package startup

import (
	"log"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("../../")
	v.SetConfigName("dev")
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
	log.Printf("Load configuration profile %v successful", serverConfig.Profile)
	log.Printf("Welcome to %v version %v ", serverConfig.Name, serverConfig.Version)
	log.Println("================================")
}
