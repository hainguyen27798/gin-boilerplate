package setting

import (
	setting2 "github.com/hainguyen27798/gin-boilerplate/pkg/setting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppMode_String(t *testing.T) {
	t.Run("should return correct string for ProdMode", func(t *testing.T) {
		mode := setting2.ProdMode
		assert.Equal(t, "prod", string(mode))
	})

	t.Run("should return correct string for DevMode", func(t *testing.T) {
		mode := setting2.DevMode
		assert.Equal(t, "dev", string(mode))
	})
}

func TestAppMode_Constants(t *testing.T) {
	t.Run("should have correct constant values", func(t *testing.T) {
		assert.NotEqual(t, setting2.ProdMode, setting2.DevMode)
		assert.Equal(t, setting2.AppMode("prod"), setting2.ProdMode)
		assert.Equal(t, setting2.AppMode("dev"), setting2.DevMode)
	})
}

func TestAppMode_TypeConversion(t *testing.T) {
	t.Run("should allow conversion from string", func(t *testing.T) {
		mode := setting2.AppMode("prod")
		assert.Equal(t, setting2.ProdMode, mode)

		mode = setting2.AppMode("dev")
		assert.Equal(t, setting2.DevMode, mode)
	})

	t.Run("should handle empty string conversion", func(t *testing.T) {
		mode := setting2.AppMode("")
		assert.NotEqual(t, setting2.ProdMode, mode)
		assert.NotEqual(t, setting2.DevMode, mode)
		assert.Equal(t, setting2.AppMode(""), mode)
	})
}
