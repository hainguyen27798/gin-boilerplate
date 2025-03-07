package response

import (
	"errors"
	"net/http"
	"testing"

	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	appErr := errors.New("application error")
	serviceErr := errors.New("service error")

	err := response.NewError(appErr, serviceErr)

	assert.NotNil(t, err)
	assert.Equal(t, appErr.Error(), err.AppErr())
	assert.Equal(t, serviceErr.Error(), err.ServiceErr())
}

func TestError_Error(t *testing.T) {
	appErr := errors.New("application error")
	serviceErr := errors.New("service error")

	err := response.NewError(appErr, serviceErr)

	assert.Contains(t, err.Error(), "application error")
	assert.Contains(t, err.Error(), "service error")
}

func TestError_Code(t *testing.T) {
	testCases := []struct {
		name     string
		appErr   error
		expected int
	}{
		{"BadRequest", response.ErrBadRequest, http.StatusBadRequest},
		{"JWTInternalError", response.ErrJWTInternalError, http.StatusInternalServerError},
		{"NotFound", response.ErrNotFound, http.StatusNotFound},
		{"Unauthorized", response.ErrUnauthorized, http.StatusUnauthorized},
		{"InternalError", response.ErrInternalError, http.StatusInternalServerError},
		{"InvalidToken", response.ErrInvalidToken, http.StatusUnauthorized},
		{"ExpiredToken", response.ErrExpiredToken, http.StatusUnauthorized},
		{"StolenToken", response.ErrStolenToken, http.StatusUnauthorized},
		{"InvalidObjectID", response.ErrInvalidObjectID, http.StatusBadRequest},
		{"Validation", response.ErrValidation, http.StatusBadRequest},
		{"DefaultError", errors.New("unknown error"), http.StatusInternalServerError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := response.NewError(tc.appErr, errors.New("service error"))
			assert.Equal(t, tc.expected, err.Code())
		})
	}
}

func TestError_AppErr(t *testing.T) {
	appErr := response.ErrNotFound
	serviceErr := errors.New("service error")

	err := response.NewError(appErr, serviceErr)

	assert.Equal(t, appErr.Error(), err.AppErr())
}

func TestError_ServiceErr(t *testing.T) {
	appErr := response.ErrBadRequest
	serviceErr := errors.New("service error")

	err := response.NewError(appErr, serviceErr)

	assert.Equal(t, serviceErr.Error(), err.ServiceErr())
}
