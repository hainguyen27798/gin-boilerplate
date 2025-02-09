package users

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/internal/module/users"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/internal/initialize"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment(t *testing.T) (*users.UserController, users.UserService) {
	initialize.RegisterValidations()
	helpers.Must(os.Setenv("MODE", "test"))
	defer func() {
		helpers.Must(os.Unsetenv("MODE"))
	}()

	initialize.LoadConfig("../../../../configs/")
	initialize.InitLogger()
	initialize.InitDatabase()

	repo := users.NewUserRepository(global.MongoDB.DB)
	userService := users.NewUserService(repo)
	userController := users.NewUserController(userService)

	collection := global.MongoDB.DB.Collection(users.UserModel{}.CollectionName())
	if err := collection.Drop(context.Background()); err != nil {
		t.Fatal(err)
	}

	return userController, userService
}

func TestUserController_CreateUser(t *testing.T) {
	initialize.RegisterValidations()
	controller, _ := setupTestEnvironment(t)

	defer func() {
		if global.MongoDB != nil {
			_ = global.MongoDB.Disconnect(context.Background())
		}
	}()

	t.Run("successful user creation", func(t *testing.T) {
		userData := users.CreateUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@example.com",
			Password:  "StrongPass123!",
		}
		jsonData, _ := json.Marshal(userData)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.CreateUser(c)

		var res response.TDataResponse
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, response.CreatedSuccess, res.Code)
		assert.Equal(t, response.CodeMsg[response.CreatedSuccess], res.Message)

		responseData := res.Data.(map[string]interface{})
		assert.Equal(t, userData.Email, responseData["email"])
		assert.Equal(t, userData.FirstName, responseData["first_name"])
		assert.Equal(t, userData.LastName, responseData["last_name"])
		assert.Empty(t, responseData["password"])
		assert.Empty(t, responseData["verification_code"])
		assert.False(t, responseData["verified"].(bool))
	})

	t.Run("invalid user data", func(t *testing.T) {
		userData := users.CreateUserDto{
			FirstName: "",
			LastName:  "Doe",
			Email:     "invalid-email",
			Password:  "weak",
		}
		jsonData, _ := json.Marshal(userData)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.CreateUser(c)

		var res response.TErrResponse
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, response.ErrCodeParamInvalid, res.Code)
		assert.Equal(t, response.CodeMsg[response.ErrCodeParamInvalid], res.Message)

		errors := res.Errors.([]interface{})
		assert.Contains(t, errors, "Field FirstName is required")
		assert.Contains(t, errors, "Fields Email is invalid due to email")
		assert.Contains(t, errors, "Fields Password is invalid due to strongPassword")
	})
}

func TestUserController_GetUserByID(t *testing.T) {
	controller, service := setupTestEnvironment(t)

	t.Run("get existing user", func(t *testing.T) {
		userData := users.CreateUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@example.com",
			Password:  "StrongPass123!",
		}
		ctx := context.Background()
		userCreated, err := service.CreateUser(ctx, &userData)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: userCreated.ID}}
		c.Request = req

		controller.GetUserByID(c)

		var res response.TDataResponse
		err = json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, response.CodeSuccess, res.Code)
		assert.Equal(t, response.CodeMsg[response.CodeSuccess], res.Message)

		responseData := res.Data.(map[string]interface{})
		assert.Equal(t, userData.Email, responseData["email"])
		assert.Equal(t, userData.FirstName, responseData["first_name"])
		assert.Equal(t, userData.LastName, responseData["last_name"])
	})

	t.Run("invalid user id format", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: "invalid-id"}}
		c.Request = req

		controller.GetUserByID(c)

		var res response.TResponse
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, response.ErrCodeQueryParamInvalid, res.Code)
		assert.Equal(t, response.CodeMsg[response.ErrCodeQueryParamInvalid], res.Message)
	})
}

func TestUserController_UpdateUser(t *testing.T) {
	controller, service := setupTestEnvironment(t)

	t.Run("successful update", func(t *testing.T) {
		// Create test user
		userData := users.CreateUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@example.com",
			Password:  "StrongPass123!",
		}
		ctx := context.Background()
		userCreated, err := service.CreateUser(ctx, &userData)
		assert.NoError(t, err)

		// Update user
		updateData := users.UpdateUserDto{
			FirstName: "Updated John",
			LastName:  "Updated Doe",
		}
		jsonData, _ := json.Marshal(updateData)

		req := httptest.NewRequest(http.MethodPut, "/users/:id", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: userCreated.ID}}
		c.Request = req

		controller.UpdateUser(c)

		var res response.TDataResponse
		err = json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, response.CodeSuccess, res.Code)
		assert.Equal(t, response.CodeMsg[response.CodeSuccess], res.Message)

		responseData := res.Data.(map[string]interface{})
		assert.Equal(t, updateData.FirstName, responseData["first_name"])
		assert.Equal(t, updateData.LastName, responseData["last_name"])
	})
}

func TestUserController_DeleteUser(t *testing.T) {
	controller, service := setupTestEnvironment(t)

	t.Run("successful deletion", func(t *testing.T) {
		userData := users.CreateUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@example.com",
			Password:  "StrongPass123!",
		}
		ctx := context.Background()
		userCreated, err := service.CreateUser(ctx, &userData)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "id", Value: userCreated.ID}}
		c.Request = req

		controller.DeleteUser(c)

		var res response.TResponse
		err = json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, response.CodeSuccess, res.Code)
		assert.Equal(t, response.CodeMsg[response.CodeSuccess], res.Message)
	})
}

func TestUserController_GetUserByEmail(t *testing.T) {
	controller, service := setupTestEnvironment(t)

	t.Run("successful email lookup", func(t *testing.T) {
		userData := users.CreateUserDto{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "test@example.com",
			Password:  "StrongPass123!",
		}
		ctx := context.Background()
		_, err := service.CreateUser(ctx, &userData)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req
		q := c.Request.URL.Query()
		q.Add("email", userData.Email)
		c.Request.URL.RawQuery = q.Encode()

		controller.GetUserByEmail(c)

		var res response.TDataResponse
		err = json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, response.CodeSuccess, res.Code)
		assert.Equal(t, response.CodeMsg[response.CodeSuccess], res.Message)

		responseData := res.Data.(map[string]interface{})
		assert.Equal(t, userData.Email, responseData["email"])
		assert.Equal(t, userData.FirstName, responseData["first_name"])
		assert.Equal(t, userData.LastName, responseData["last_name"])
	})

	t.Run("missing email parameter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		controller.GetUserByEmail(c)

		var res response.TResponse
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, response.ErrCodeParamInvalid, res.Code)
		assert.Equal(t, response.CodeMsg[response.ErrCodeParamInvalid], res.Message)
	})
}
