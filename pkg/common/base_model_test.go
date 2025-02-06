package common

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBaseModel_BeforeCreate(t *testing.T) {
	existingID := bson.NewObjectID()

	t.Run("should set ID and timestamps when empty", func(t *testing.T) {
		model := &BaseModel{}
		model.BeforeCreate()

		assert.NotEmpty(t, model.ID)
		assert.False(t, model.CreatedAt.IsZero())
		assert.False(t, model.UpdatedAt.IsZero())
		assert.Equal(t, model.CreatedAt.UTC().Location(), time.UTC)
		assert.Equal(t, model.UpdatedAt.UTC().Location(), time.UTC)
	})

	t.Run("should preserve existing ID and CreatedAt", func(t *testing.T) {
		existingTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

		model := &BaseModel{
			ID:        existingID,
			CreatedAt: existingTime,
		}
		model.BeforeCreate()

		assert.Equal(t, existingID, model.ID)
		assert.Equal(t, existingTime, model.CreatedAt)
		assert.False(t, model.UpdatedAt.IsZero())
		assert.True(t, model.UpdatedAt.After(existingTime))
	})
}

func TestBaseModel_BeforeUpdate(t *testing.T) {
	t.Run("should update UpdatedAt timestamp", func(t *testing.T) {
		model := &BaseModel{
			ID:        bson.NewObjectID(),
			CreatedAt: time.Now().Add(-24 * time.Hour),
			UpdatedAt: time.Now().Add(-24 * time.Hour),
		}

		oldUpdatedAt := model.UpdatedAt
		time.Sleep(time.Millisecond)
		model.BeforeUpdate()

		assert.True(t, model.UpdatedAt.After(oldUpdatedAt))
		assert.Equal(t, time.UTC, model.UpdatedAt.Location())
	})

	t.Run("should preserve other fields", func(t *testing.T) {
		originalID := bson.NewObjectID()
		originalCreatedAt := time.Now().Add(-24 * time.Hour)

		model := &BaseModel{
			ID:        originalID,
			CreatedAt: originalCreatedAt,
		}

		model.BeforeUpdate()

		assert.Equal(t, originalID.Hex(), model.ID.Hex())
		assert.Equal(t, originalCreatedAt, model.CreatedAt)
	})
}
