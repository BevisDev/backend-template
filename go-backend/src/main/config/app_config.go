package config

var AppConfig Config

type Config struct {
	// Server
	ServerConfig struct {
		Name           string   `mapstructure:"name"`
		Profile        string   `mapstructure:"profile"`
		TrustedProxies []string `mapstructure:"trustedProxies"`
		Port           string   `mapstructure:"port"`
		Version        string   `mapstructure:"version"`
		Timezone       string   `mapstructure:"timezone"`
		ClientTimeout  int      `mapstructure:"clientTimeout"`
		ServerTimeout  int      `mapstructure:"serverTimeout"`
	} `mapstructure:"server"`

	// Logger
	LoggerConfig struct {
		IsSplit    bool   `mapstructure:"isSplit"`
		CronTime   string `mapstructure:"cronTime"`
		LogAppDir  string `mapstructure:"logAppDir"`
		LogRRDir   string `mapstructure:"logRRDir"`
		LogExtDir  string `mapstructure:"logExtDir"`
		MaxSize    int    `mapstructure:"maxSize"`
		MaxBackups int    `mapstructure:"maxBackups"`
		MaxAge     int    `mapstructure:"maxAge"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"logger"`

	// Databases
	Databases []Database `mapstructure:"databases"`

	// Redis
	RedisConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Index    int    `mapstructure:"index"`
	} `mapstructure:"redis"`
}

type Database struct {
	Driver          string   `mapstructure:"driver"`
	Host            string   `mapstructure:"host"`
	Port            int      `mapstructure:"port"`
	Username        string   `mapstructure:"username"`
	Password        string   `mapstructure:"password"`
	TimeoutSec      int      `mapstructure:"timeoutSec"`
	MaxOpenConns    int      `mapstructure:"maxOpenConns"`
	MaxIdleConns    int      `mapstructure:"maxIdleConns"`
	ConnMaxLifeTime int      `mapstructure:"connMaxLifeTime"`
	Schema          []string `mapstructure:"schema"`
}
