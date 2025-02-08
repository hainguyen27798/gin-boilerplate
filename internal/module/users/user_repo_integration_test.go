package users

import (
	"context"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"os"
	"testing"
)

func TestUserRepository_Integration(t *testing.T) {
	// Setup test config
	helpers.Must(os.Setenv("MODE", "test"))
	defer func() {
		helpers.Must(os.Unsetenv("MODE"))
	}()

	initialize.LoadConfig("../../../configs/")
	initialize.InitLogger()

	ctx := context.Background()

	// Initialize the database using the shared InitDatabase() method.
	initialize.InitDatabase()
	defer func() {
		if global.MongoDB != nil {
			err := global.MongoDB.Disconnect(ctx)
			if err != nil {
				panic(err)
			}
		}
	}()

	// Clean up the users collection to start fresh.
	repo := NewUserRepository(global.MongoDB.DB)
	collection := global.MongoDB.DB.Collection(UserModel{}.CollectionName())
	if err := collection.Drop(ctx); err != nil {
		assert.NoError(t, err)
	}

	// Prepare a test user.
	hexId, _ := bson.ObjectIDFromHex("67a4f57c39b9abb0dbabd5b0")
	user := &UserModel{
		BaseModel: common.BaseModel{
			ID: hexId,
		},
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "StrongP@ss123!",
		Image:     "test.jpg",
	}

	t.Run("Create user", func(t *testing.T) {
		newUser, err := repo.Create(ctx, user)
		assert.NoError(t, err)
		assert.NotNil(t, newUser)
		assert.Equal(t, user.Email, newUser.Email)
		assert.Equal(t, user.FirstName, newUser.FirstName)
		assert.Equal(t, user.LastName, newUser.LastName)
		assert.Equal(t, user.Image, newUser.Image)
		assert.Equal(t, user.Password, newUser.Password)
		assert.NotEmpty(t, newUser.CreatedAt)
		assert.NotEmpty(t, newUser.UpdatedAt)
	})

	t.Run("Find user by email", func(t *testing.T) {
		found, err := repo.FindByEmail(ctx, user.Email)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, found.Email)
	})

	t.Run("Find user by ID", func(t *testing.T) {
		found, err := repo.FindByID(ctx, user.ID.Hex())

		assert.NoError(t, err)
		assert.Equal(t, user.ID, found.ID)
	})

	t.Run("Update fistName user", func(t *testing.T) {
		firstName := "Jane"

		// Update timestamp.
		userUpdated, err := repo.Update(ctx, user.ID.Hex(), bson.D{
			{
				"$set",
				bson.D{
					{"first_name", firstName},
				},
			},
		})
		assert.NoError(t, err)
		updated := userUpdated.ToDto()
		assert.Equal(t, firstName, updated.FirstName)
		assert.NotEmpty(t, updated.UpdatedAt)
	})

	t.Run("Delete user", func(t *testing.T) {
		err := repo.Delete(ctx, user.ID.Hex())
		assert.NoError(t, err)

		deleted, err := repo.FindByID(ctx, user.ID.Hex())
		assert.Error(t, err)
		assert.Nil(t, deleted)
	})
}
