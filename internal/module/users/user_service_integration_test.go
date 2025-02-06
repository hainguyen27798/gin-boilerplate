package users

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
)

func TestUserService_Integration(t *testing.T) {
	// Setup environment and initialize configuration, logger, and database.
	helpers.Must(os.Setenv("MODE", "test"))
	defer func() {
		helpers.Must(os.Unsetenv("MODE"))
	}()
	initialize.LoadConfig("../../../configs/")
	initialize.InitLogger()
	initialize.InitDatabase()

	ctx := context.Background()
	defer func() {
		if global.MongoDB != nil {
			if err := global.MongoDB.Disconnect(ctx); err != nil {
				panic(err)
			}
		}
	}()

	// Ensure a clean state.
	collection := global.MongoDB.DB.Collection(UserModel{}.CollectionName())
	if err := collection.Drop(ctx); err != nil {
		assert.NoError(t, err)
	}

	// Create a new UserService instance using the repository.
	repo := NewUserRepository(global.MongoDB.DB)
	service := NewUserService(repo)

	// Prepare a new user DTO.
	createDTO := &CreateUserDto{
		Email:     "testservice@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "StrongP@ss123!",
		Image:     "test.jpg",
	}

	var createdUserID string

	t.Run("Create user", func(t *testing.T) {
		err := service.CreateUser(ctx, createDTO)
		assert.NoError(t, err)
	})

	t.Run("Get user by email", func(t *testing.T) {
		userDto, err := service.GetUserByEmail(ctx, createDTO.Email)
		assert.NoError(t, err)
		assert.Equal(t, createDTO.Email, userDto.Email)
		assert.Equal(t, createDTO.FirstName, userDto.FirstName)
		assert.Equal(t, createDTO.LastName, userDto.LastName)
		createdUserID = userDto.ID
	})

	t.Run("Get user by ID", func(t *testing.T) {
		userDto, err := service.GetUserByID(ctx, createdUserID)
		assert.NoError(t, err)
		assert.Equal(t, createDTO.Email, userDto.Email)
	})

	t.Run("Update user", func(t *testing.T) {
		// Prepare an update DTO with the modified first name.
		updateDTO := &UpdateUserDto{
			FirstName: "Jane",
		}

		err := service.UpdateUser(ctx, createdUserID, updateDTO)
		assert.NoError(t, err)

		// Retrieve updated user.
		userDto, err := service.GetUserByID(ctx, createdUserID)
		assert.NoError(t, err)
		assert.Equal(t, "Jane", userDto.FirstName)
	})

	t.Run("Delete user", func(t *testing.T) {
		err := service.DeleteUser(ctx, createdUserID)
		assert.NoError(t, err)

		// Attempt to fetch the deleted user.
		userDto, err := service.GetUserByID(ctx, createdUserID)
		assert.Error(t, err)
		assert.Nil(t, userDto)
	})
}
