package setting

// Config represents the application's configuration structure for server settings.
type Config struct {
	Server  ServerSettings  `mapstructure:"server_config"`
	Logger  LoggerSettings  `mapstructure:"logger_config"`
	MongoDB MongoDBSettings `mapstructure:"mongo_config"`
}

// ServerSettings defines the configuration settings for a server, including the port it operates on.
type ServerSettings struct {
	Port string `mapstructure:"port"`
}

// LoggerSettings is a configuration structure for setting up logging behavior and file management.
type LoggerSettings struct {
	FileName   string `mapstructure:"file_name"`
	Level      string `mapstructure:"log_level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// MongoDBSettings defines the configuration settings for mongoDB
type MongoDBSettings struct {
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Database         string `mapstructure:"database"`
	MaxPoolSize      uint64 `mapstructure:"max_pool_size"`
	EnableLog        bool   `mapstructure:"enable_log"`
	DirectConnection bool   `mapstructure:"direct_connection"`
}
