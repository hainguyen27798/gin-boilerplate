package initialize

import (
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"github.com/spf13/viper"
	"os"
)

// LoadConfig initializes and reads the application configuration using the viper library and unmarshals it into AppConfig.
func LoadConfig() {
	mode := os.Getenv("MODE")

	// Set App Mode
	switch setting.AppMode(mode) {
	case setting.DevMode, setting.ProdMode:
		global.AppMode = setting.AppMode(mode)
		break
	default:
		global.AppMode = setting.DevMode
	}

	viper.AddConfigPath("./configs/")
	viper.SetConfigName(string(global.AppMode))
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&global.AppConfig)
	if err != nil {
		panic(err)
	}
}
