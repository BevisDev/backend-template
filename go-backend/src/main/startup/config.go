package startup

import (
	"log"

	"github.com/BevisDev/backend-template/src/main/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("../../")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	// read config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error can not read configuration %v", err)
	}

	if err := viper.Unmarshal(&global.AppConfig); err != nil {
		log.Fatalf("Error decode config into struct, %v", err)
	}

	log.Printf("Load configuration profile %v successful", global.AppConfig.ServerConfig.Profile)
	log.Printf("Welcome %v version %v ", global.AppConfig.ServerConfig.Name, global.AppConfig.ServerConfig.Version)
}
