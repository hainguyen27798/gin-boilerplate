package setting

// Config represents the application's configuration structure for server settings.
type Config struct {
	Server ServerConfig `mapstructure:"server_config"`
}

// ServerConfig holds configuration settings for the server, including the port configuration.
type ServerConfig struct {
	Port string `mapstructure:"port"`
}
