package config

type AppConfig struct {
	// Server
	ServerConfig struct {
		Name           string   `mapstructure:"name"`
		Profile        string   `mapstructure:"profile"`
		Port           string   `mapstructure:"port"`
		Version        string   `mapstructure:"version"`
		TrustedProxies []string `mapstructure:"trustedProxies"`
	} `mapstructure:"server"`

	// Databases
	DatabasesConfig struct {
		MySqlConfig      DatabaseProps `mapstructure:"mysql"`
		PostgreSqlConfig DatabaseProps `mapstructure:"postgresql"`
	} `mapstructure:"databases"`

	// LOGGER
	LoggerConfig struct {
		LogLevel   string `mapstructure:"logLevel"`
		Filename   string `mapstructure:"filename"`
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
