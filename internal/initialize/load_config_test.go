package initialize

import (
	"os"
	"testing"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("should load dev config when MODE is dev", func(t *testing.T) {
		helpers.Must(os.Setenv("MODE", "dev"))
		defer func() {
			helpers.Must(os.Unsetenv("MODE"))
		}()

		LoadConfig("../../configs/")

		assert.Equal(t, setting.DevMode, global.AppMode)
		assert.NotNil(t, global.AppConfig)
	})

	t.Run("should load prod config when MODE is prod", func(t *testing.T) {
		helpers.Must(os.Setenv("MODE", "prod"))
		defer func() {
			helpers.Must(os.Unsetenv("MODE"))
		}()

		LoadConfig("../../configs/")

		assert.Equal(t, setting.ProdMode, global.AppMode)
		assert.NotNil(t, global.AppConfig)
	})

	t.Run("should default to dev mode when MODE is invalid", func(t *testing.T) {
		helpers.Must(os.Setenv("MODE", "invalid"))
		defer func() {
			helpers.Must(os.Unsetenv("MODE"))
		}()

		LoadConfig("../../configs/")

		assert.Equal(t, setting.DevMode, global.AppMode)
		assert.NotNil(t, global.AppConfig)
	})

	t.Run("should default to dev mode when MODE is empty", func(t *testing.T) {
		helpers.Must(os.Unsetenv("MODE"))

		LoadConfig("../../configs/")

		assert.Equal(t, setting.DevMode, global.AppMode)
		assert.NotNil(t, global.AppConfig)
	})
}
