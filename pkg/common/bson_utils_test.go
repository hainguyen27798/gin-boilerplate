package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidObjectID(t *testing.T) {
	t.Run("should return true for valid ObjectID hex string", func(t *testing.T) {
		validID := "507f1f77bcf86cd799439011"
		assert.True(t, IsValidObjectID(validID))
	})

	t.Run("should return false for invalid hex string", func(t *testing.T) {
		invalidID := "invalid-hex-string"
		assert.False(t, IsValidObjectID(invalidID))
	})

	t.Run("should return false for empty string", func(t *testing.T) {
		assert.False(t, IsValidObjectID(""))
	})

	t.Run("should return false for short hex string", func(t *testing.T) {
		shortID := "507f1f"
		assert.False(t, IsValidObjectID(shortID))
	})

	t.Run("should return false for long hex string", func(t *testing.T) {
		longID := "507f1f77bcf86cd799439011507f1f77bcf86cd799439011"
		assert.False(t, IsValidObjectID(longID))
	})
}
