package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
	"github.com/stretchr/testify/assert"
)

func setupGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestOkResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, w := setupGinContext()

	data := map[string]string{"key": "value"}
	message := "Success"

	response.OkResponse(c, message, data)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), message)
	assert.Contains(t, w.Body.String(), "\"key\":\"value\"")
}

func TestCreatedResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, w := setupGinContext()

	data := map[string]string{"id": "123"}
	message := "Resource created"

	response.CreatedResponse(c, message, data)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), message)
	assert.Contains(t, w.Body.String(), "\"id\":\"123\"")
}

type MockValidationErrors struct {
	validator.FieldError
}

func (m MockValidationErrors) Field() string {
	return "username"
}

func (m MockValidationErrors) Tag() string {
	return "required"
}

func (m MockValidationErrors) Error() string {
	return "validation error"
}

func TestValidateErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("with validation errors", func(t *testing.T) {
		c, w := setupGinContext()

		// Create mock validation errors
		var validationErrors validator.ValidationErrors
		validationErrors = append(validationErrors, MockValidationErrors{})

		response.ValidateErrorResponse(c, validationErrors)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "validation error")
		assert.Contains(t, w.Body.String(), "Field username is required")
	})

	t.Run("with regular error", func(t *testing.T) {
		c, w := setupGinContext()

		err := errors.New("some error")
		response.ValidateErrorResponse(c, err)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "validation error")
	})
}

type MockStruct struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Age     int    `json:"age" validate:"min=18"`
	Website string `json:"website" validate:"url"`
}

func TestValidationErrorsToJSON(t *testing.T) {
	// This is tested indirectly through ValidateErrorResponse
	// since validationErrorsToJSON is a private function

	gin.SetMode(gin.TestMode)
	c, w := setupGinContext()

	validate := validator.New()
	testStruct := MockStruct{
		Name:    "",
		Email:   "invalid-email",
		Age:     15,
		Website: "not-a-url",
	}

	err := validate.Struct(testStruct)
	response.ValidateErrorResponse(c, err)

	// Check that validation errors are properly formatted
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Field Name is required")
	assert.Contains(t, w.Body.String(), "Fields Email is invalid due to email")
	assert.Contains(t, w.Body.String(), "Fields Age is invalid due to min")
	assert.Contains(t, w.Body.String(), "Fields Website is invalid due to url")
}

func TestErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name           string
		appErr         error
		serviceErr     error
		expectedStatus int
		expectedBody   []string
	}{
		{
			name:           "bad request error",
			appErr:         response.ErrBadRequest,
			serviceErr:     errors.New("invalid parameters"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []string{"bad request", "invalid parameters"},
		},
		{
			name:           "not found error",
			appErr:         response.ErrNotFound,
			serviceErr:     errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   []string{"resource not found", "user not found"},
		},
		{
			name:           "internal server error",
			appErr:         response.ErrInternalError,
			serviceErr:     errors.New("database connection failed"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []string{"internal error", "database connection failed"},
		},
		{
			name:           "unauthorized error",
			appErr:         response.ErrUnauthorized,
			serviceErr:     errors.New("invalid credentials"),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   []string{"unauthorized", "invalid credentials"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			err := response.NewError(tc.appErr, tc.serviceErr)
			response.ErrorResponse(c, err)

			assert.Equal(t, tc.expectedStatus, w.Code)

			for _, expectedStr := range tc.expectedBody {
				assert.Contains(t, w.Body.String(), expectedStr)
			}
		})
	}
}

func TestMessageResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name           string
		appErr         error
		expectedStatus int
		expectedMsg    string
	}{
		{
			name:           "bad request error",
			appErr:         response.ErrBadRequest,
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "bad request",
		},
		{
			name:           "not found error",
			appErr:         response.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedMsg:    "resource not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			response.ErrorResponse(c, response.NewError(tc.appErr, nil))

			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tc.expectedMsg)
		})
	}
}
