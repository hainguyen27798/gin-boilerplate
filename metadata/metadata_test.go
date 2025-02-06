package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadataVariables(t *testing.T) {
	t.Run("should have default values", func(t *testing.T) {
		assert.Equal(t, "v-local", Version)
		assert.Equal(t, "unknown", Commit)
		assert.Equal(t, "unknown", BuildDate)
	})

	t.Run("should allow overriding values", func(t *testing.T) {
		oldVersion := Version
		oldCommit := Commit
		oldBuildDate := BuildDate

		Version = "v1.0.0"
		Commit = "abc123"
		BuildDate = "2024-01-01"

		assert.Equal(t, "v1.0.0", Version)
		assert.Equal(t, "abc123", Commit)
		assert.Equal(t, "2024-01-01", BuildDate)

		// Restore original values
		Version = oldVersion
		Commit = oldCommit
		BuildDate = oldBuildDate
	})
}
