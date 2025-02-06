package initialize

import (
	"context"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"os"
	"testing"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/stretchr/testify/assert"
)

func TestInitDatabase(t *testing.T) {
	t.Run("should successfully initialize database with valid config", func(t *testing.T) {
		// Setup test config
		helpers.Must(os.Setenv("MODE", "test"))
		defer func() {
			helpers.Must(os.Unsetenv("MODE"))
		}()

		LoadConfig("../../configs/")
		InitLogger()

		// Test execution
		InitDatabase()

		defer func() {
			if global.MongoDB != nil {
				helpers.Must(global.MongoDB.Disconnect(context.Background()))
			}
		}()

		// Assertions
		assert.NotNil(t, global.MongoDB)
	})
}
