package users

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserDto_Validate(t *testing.T) {
	initialize.RegisterValidations()

	t.Run("valid user dto", func(t *testing.T) {
		dto := CreateUserDto{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "StrongPass123!",
		}
		err := dto.Validate()
		if ok := assert.NoError(t, err); !ok {
			panic(err)
		}
	})

	t.Run("missing required fields", func(t *testing.T) {
		dto := CreateUserDto{}
		err := dto.Validate()
		assert.Error(t, err)
	})

	t.Run("invalid email format", func(t *testing.T) {
		dto := CreateUserDto{
			Email:     "invalid-email",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "StrongPass123!",
		}
		err := dto.Validate()
		var validationErrs validator.ValidationErrors
		assert.Error(t, err)
		errors.As(err, &validationErrs)
		assert.Equal(t, 1, len(validationErrs))
		assert.Equal(t, "email", validationErrs[0].Tag())
	})

	t.Run("invalid image url", func(t *testing.T) {
		dto := CreateUserDto{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "StrongPass123!",
			Image:     "invalid-url",
		}
		err := dto.Validate()
		var validationErrs validator.ValidationErrors
		assert.Error(t, err)
		errors.As(err, &validationErrs)
		assert.Equal(t, 1, len(validationErrs))
		assert.Equal(t, "url", validationErrs[0].Tag())
	})

	t.Run("invalid strong password", func(t *testing.T) {
		dto := CreateUserDto{
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "weak",
			Image:     "https://example.com/image.jpg",
		}
		err := dto.Validate()
		var validationErrs validator.ValidationErrors
		assert.Error(t, err)
		errors.As(err, &validationErrs)
		assert.Equal(t, 1, len(validationErrs))
		assert.Equal(t, "strongPassword", validationErrs[0].Tag())
	})
}

func TestUpdateUserDto(t *testing.T) {
	t.Run("should allow partial updates", func(t *testing.T) {
		dto := UpdateUserDto{
			FirstName: "John",
		}
		err := dto.Validate()
		assert.NoError(t, err)
		assert.Equal(t, "John", dto.FirstName)
		assert.Empty(t, dto.LastName)
		assert.Empty(t, dto.Image)
	})
}

func TestUserDto(t *testing.T) {
	t.Run("should contain all required fields", func(t *testing.T) {
		dto := UserDto{
			BaseDto: common.BaseDto{
				ID: "123",
			},
			Email:     "test@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Image:     "https://example.com/image.jpg",
			Verified:  true,
		}

		assert.Equal(t, "123", dto.ID)
		assert.Equal(t, "test@example.com", dto.Email)
		assert.Equal(t, "John", dto.FirstName)
		assert.Equal(t, "Doe", dto.LastName)
		assert.Equal(t, "https://example.com/image.jpg", dto.Image)
		assert.True(t, dto.Verified)
	})
}
