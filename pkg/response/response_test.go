package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestMessageResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("should return correct status code for success range", func(t *testing.T) {
		MessageResponse(c, CodeSuccess)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), CodeMsg[CodeSuccess])
	})

	t.Run("should return correct status code for error range", func(t *testing.T) {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		MessageResponse(c, ErrBadRequest)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), CodeMsg[ErrBadRequest])
	})
}

func TestOkResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("should include data in response", func(t *testing.T) {
		testData := map[string]string{"test": "value"}
		OkResponse(c, CodeSuccess, testData)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"test":"value"`)
	})
}

func TestCreatedResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("should return created status with data", func(t *testing.T) {
		testData := map[string]int{"id": 1}
		CreatedResponse(c, CreatedSuccess, testData)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"id":1`)
	})
}

func TestValidateErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("should handle validation errors", func(t *testing.T) {
		validate := validator.New()
		type TestStruct struct {
			Field string `validate:"required"`
		}
		var test TestStruct
		err := validate.Struct(test)
		ValidateErrorResponse(c, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Field Field is required")
	})

	t.Run("should handle non-validation errors", func(t *testing.T) {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		err := errors.New("random error")
		ValidateErrorResponse(c, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NotContains(t, w.Body.String(), "random error")
	})
}

func TestGetHTTPCode(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected int
	}{
		{"success code 20000-20099", 20050, http.StatusOK},
		{"success code 20100-20199", 20150, http.StatusOK},
		{"error code 40000-49999", 40001, http.StatusBadRequest},
		{"internal error code", 50000, http.StatusInternalServerError},
		{"unknown code below range", 10000, http.StatusInternalServerError},
		{"unknown code above range", 60000, http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getHTTPCode(tt.code)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNotFoundException(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	t.Run("should return not found status and abort", func(t *testing.T) {
		NotFoundException(c, ErrNotFound)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), CodeMsg[ErrNotFound])
	})
}
