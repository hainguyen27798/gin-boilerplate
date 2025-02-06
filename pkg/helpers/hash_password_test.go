package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("should successfully hash password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, err := HashPassword(password)

		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEqual(t, password, hash)
	})

	t.Run("should generate different hashes for same password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash1, _ := HashPassword(password)
		hash2, _ := HashPassword(password)

		assert.NotEqual(t, hash1, hash2)
	})

	t.Run("should handle empty password", func(t *testing.T) {
		hash, err := HashPassword("")

		assert.NoError(t, err)
		assert.NotEmpty(t, hash)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("should verify correct password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, _ := HashPassword(password)

		assert.True(t, CheckPasswordHash(password, hash))
	})

	t.Run("should reject incorrect password", func(t *testing.T) {
		password := "mySecurePassword123"
		hash, _ := HashPassword(password)

		assert.False(t, CheckPasswordHash("wrongPassword", hash))
	})

	t.Run("should handle empty password and hash", func(t *testing.T) {
		assert.False(t, CheckPasswordHash("", ""))
	})

	t.Run("should reject invalid hash format", func(t *testing.T) {
		assert.False(t, CheckPasswordHash("password", "invalidhash"))
	})
}
