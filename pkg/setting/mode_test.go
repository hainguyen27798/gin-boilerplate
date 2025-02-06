package setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppMode_String(t *testing.T) {
	t.Run("should return correct string for ProdMode", func(t *testing.T) {
		mode := ProdMode
		assert.Equal(t, "prod", string(mode))
	})

	t.Run("should return correct string for DevMode", func(t *testing.T) {
		mode := DevMode
		assert.Equal(t, "dev", string(mode))
	})
}

func TestAppMode_Constants(t *testing.T) {
	t.Run("should have correct constant values", func(t *testing.T) {
		assert.NotEqual(t, ProdMode, DevMode)
		assert.Equal(t, AppMode("prod"), ProdMode)
		assert.Equal(t, AppMode("dev"), DevMode)
	})
}

func TestAppMode_TypeConversion(t *testing.T) {
	t.Run("should allow conversion from string", func(t *testing.T) {
		mode := AppMode("prod")
		assert.Equal(t, ProdMode, mode)

		mode = AppMode("dev")
		assert.Equal(t, DevMode, mode)
	})

	t.Run("should handle empty string conversion", func(t *testing.T) {
		mode := AppMode("")
		assert.NotEqual(t, ProdMode, mode)
		assert.NotEqual(t, DevMode, mode)
		assert.Equal(t, AppMode(""), mode)
	})
}
