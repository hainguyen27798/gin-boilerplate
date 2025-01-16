package configs

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"SERVER_CONFIG"`
}

var AppConfig Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
