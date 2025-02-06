package setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_Structure(t *testing.T) {
	t.Run("should have all required fields", func(t *testing.T) {
		config := Config{}
		assert.NotNil(t, config.Server)
		assert.NotNil(t, config.Logger)
		assert.NotNil(t, config.MongoDB)
	})
}

func TestServerSettings_Validation(t *testing.T) {
	t.Run("should handle empty port", func(t *testing.T) {
		settings := ServerSettings{}
		assert.Empty(t, settings.Port)
	})

	t.Run("should store port value", func(t *testing.T) {
		settings := ServerSettings{Port: "8080"}
		assert.Equal(t, "8080", settings.Port)
	})
}

func TestLoggerSettings_Validation(t *testing.T) {
	t.Run("should initialize with zero values", func(t *testing.T) {
		settings := LoggerSettings{}
		assert.Empty(t, settings.FileName)
		assert.Empty(t, settings.Level)
		assert.Zero(t, settings.MaxSize)
		assert.Zero(t, settings.MaxBackups)
		assert.Zero(t, settings.MaxAge)
		assert.False(t, settings.Compress)
	})

	t.Run("should store all field values", func(t *testing.T) {
		settings := LoggerSettings{
			FileName:   "app.log",
			Level:      "info",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
			Compress:   true,
		}
		assert.Equal(t, "app.log", settings.FileName)
		assert.Equal(t, "info", settings.Level)
		assert.Equal(t, 100, settings.MaxSize)
		assert.Equal(t, 3, settings.MaxBackups)
		assert.Equal(t, 7, settings.MaxAge)
		assert.True(t, settings.Compress)
	})
}

func TestMongoDBSettings_Validation(t *testing.T) {
	t.Run("should initialize with zero values", func(t *testing.T) {
		settings := MongoDBSettings{}
		assert.Empty(t, settings.Host)
		assert.Empty(t, settings.Port)
		assert.Empty(t, settings.Username)
		assert.Empty(t, settings.Password)
		assert.Empty(t, settings.Database)
		assert.Zero(t, settings.MaxPoolSize)
		assert.False(t, settings.EnableLog)
		assert.False(t, settings.DirectConnection)
	})

	t.Run("should store all field values", func(t *testing.T) {
		settings := MongoDBSettings{
			Host:             "localhost",
			Port:             "27017",
			Username:         "admin",
			Password:         "password",
			Database:         "testdb",
			MaxPoolSize:      100,
			EnableLog:        true,
			DirectConnection: true,
		}
		assert.Equal(t, "localhost", settings.Host)
		assert.Equal(t, "27017", settings.Port)
		assert.Equal(t, "admin", settings.Username)
		assert.Equal(t, "password", settings.Password)
		assert.Equal(t, "testdb", settings.Database)
		assert.Equal(t, uint64(100), settings.MaxPoolSize)
		assert.True(t, settings.EnableLog)
		assert.True(t, settings.DirectConnection)
	})
}
