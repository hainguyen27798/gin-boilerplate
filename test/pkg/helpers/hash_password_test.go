package helpers

import (
	helpers2 "github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("should successfully hash password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, err := helpers2.HashPassword(password)

		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEqual(t, password, hash)
	})

	t.Run("should generate different hashes for same password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash1, _ := helpers2.HashPassword(password)
		hash2, _ := helpers2.HashPassword(password)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("should handle empty password", func(t *testing.T) {
		hash, err := helpers2.HashPassword("")

		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("should verify correct password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, _ := helpers2.HashPassword(password)

		assert.True(t, helpers2.CheckPasswordHash(password, hash))
	})

	t.Run("should reject incorrect password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, _ := helpers2.HashPassword(password)

		assert.False(t, helpers2.CheckPasswordHash("wrongPassword", hash))
	})

	t.Run("should handle empty password and hash", func(t *testing.T) {
		assert.False(t, helpers2.CheckPasswordHash("", ""))
	})

	t.Run("should reject invalid hash format", func(t *testing.T) {
		assert.False(t, helpers2.CheckPasswordHash("password", "invalidhash"))
	})
}
