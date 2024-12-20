package config

type AppConfig struct {
	// Server
	ServerConfig struct {
		Name           string   `mapstructure:"name"`
		Profile        string   `mapstructure:"profile"`
		TrustedProxies []string `mapstructure:"trustedProxies"`
		Port           string   `mapstructure:"port"`
		Version        string   `mapstructure:"version"`
		Timezone       string   `mapstructure:"timezone"`
	} `mapstructure:"server"`

	// Databases
	DatabasesConfig struct {
		MySqlConfig      DatabaseProps `mapstructure:"mysql"`
		PostgreSqlConfig DatabaseProps `mapstructure:"postgresql"`
	} `mapstructure:"databases"`

	// Logger
	LoggerConfig struct {
		IsSplit    string `mapstructure:"isSplit"`
		LogDir     string `mapstructure:"logDir"`
		MaxSize    int    `mapstructure:"maxSize"`
		MaxBackups int    `mapstructure:"maxBackups"`
		MaxAge     int    `mapstructure:"maxAge"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"logger"`
}

type DatabaseProps struct {
	Host     string   `mapstructure:"host"`
	Port     int      `mapstructure:"port"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	Schema   []string `mapstructure:"schema"`
}
