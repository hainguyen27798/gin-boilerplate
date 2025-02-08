package users

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"reflect"
	"testing"
	"time"

	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestUserModel_CollectionName(t *testing.T) {
	t.Run("should return correct collection name", func(t *testing.T) {
		user := UserModel{}
		assert.Equal(t, "users", user.CollectionName())
	})
}

func TestUserModel_ToDto(t *testing.T) {
	t.Run("should convert model to DTO with all fields", func(t *testing.T) {
		user := UserModel{
			BaseModel: common.BaseModel{
				ID:        bson.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Email:            "test@example.com",
			FirstName:        "John",
			LastName:         "Doe",
			Password:         "hashedpassword",
			Image:            "profile.jpg",
			Verified:         true,
			VerificationCode: "123456",
		}

		dto := user.ToDto()

		assert.Equal(t, user.ID.Hex(), dto.ID)
		assert.Equal(t, user.Email, dto.Email)
		assert.Equal(t, user.FirstName, dto.FirstName)
		assert.Equal(t, user.LastName, dto.LastName)
		assert.Equal(t, user.Image, dto.Image)
		assert.Equal(t, user.Verified, dto.Verified)
		assert.Equal(t, user.CreatedAt, dto.CreatedAt)
		assert.Equal(t, user.UpdatedAt, dto.UpdatedAt)
	})

	t.Run("should handle empty fields", func(t *testing.T) {
		user := UserModel{
			BaseModel: common.BaseModel{
				ID: bson.NewObjectID(),
			},
		}

		dto := user.ToDto()

		assert.Equal(t, user.ID.Hex(), dto.ID)
		assert.Empty(t, dto.Email)
		assert.Empty(t, dto.FirstName)
		assert.Empty(t, dto.LastName)
		assert.Empty(t, dto.Image)
		assert.Empty(t, dto.Verified)
		assert.Empty(t, dto.CreatedAt)
		assert.Empty(t, dto.UpdatedAt)
		assert.False(t, dto.Verified)
	})

	t.Run("should not expose sensitive fields", func(t *testing.T) {
		user := UserModel{
			BaseModel: common.BaseModel{
				ID: bson.NewObjectID(),
			},
			Email:            "test@example.com",
			Password:         "hashedpassword",
			VerificationCode: "123456",
		}

		dto := user.ToDto()

		assert.Equal(t, user.Email, dto.Email)

		// Iterate over the struct fields and collect their names.
		tStruct := reflect.TypeOf(*dto)
		var fieldNames []string
		for i := 0; i < tStruct.NumField(); i++ {
			field := tStruct.Field(i)
			fieldNames = append(fieldNames, field.Name)
		}

		assert.NotContains(t, fieldNames, "Password", "Password should not be exposed")
		assert.NotContains(t, fieldNames, "VerificationCode", "VerificationCode should not be exposed")
	})
}
