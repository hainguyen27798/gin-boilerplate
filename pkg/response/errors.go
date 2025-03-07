package response

import (
	"errors"
	"net/http"
)

// Generic error represents a generic error response.
var (
	ErrBadRequest       = errors.New("bad request")
	ErrJWTInternalError = errors.New("jwt internal error")
	ErrNotFound         = errors.New("resource not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrInternalError    = errors.New("internal error")
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("expired token")
	ErrStolenToken      = errors.New("stolen token")
	ErrInvalidObjectID  = errors.New("invalid object id")
	ErrValidation       = errors.New("validation error")
)

// Error represents a composite error that contains both an application-level
// error and a service-level error.
type Error struct {
	appErr     error
	serviceErr error
}

// NewError creates a new Error instance with the given application-level and
// service-level errors.
func NewError(appErr, serviceErr error) *Error {
	return &Error{
		appErr:     appErr,
		serviceErr: serviceErr,
	}
}

// Error returns a string representation of the composite error, which includes
// both the application-level error and the service-level error.
func (e *Error) Error() string {
	return errors.Join(e.appErr, e.serviceErr).Error()
}

// AppErr returns the application-level error associated with the Error.
func (e *Error) AppErr() string {
	if e.appErr == nil {
		return ""
	}
	return e.appErr.Error()
}

// ServiceErr returns the service-level error associated with the Error.
func (e *Error) ServiceErr() string {
	if e.serviceErr == nil {
		return ""
	}
	return e.serviceErr.Error()
}

// Code returns the appropriate HTTP status code based on the application-level
// error associated with the Error.
func (e *Error) Code() int {
	switch {
	case errors.Is(e.appErr, ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(e.appErr, ErrJWTInternalError):
		return http.StatusInternalServerError
	case errors.Is(e.appErr, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(e.appErr, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(e.appErr, ErrInternalError):
		return http.StatusInternalServerError
	case errors.Is(e.appErr, ErrInvalidToken):
		return http.StatusUnauthorized
	case errors.Is(e.appErr, ErrExpiredToken):
		return http.StatusUnauthorized
	case errors.Is(e.appErr, ErrStolenToken):
		return http.StatusUnauthorized
	case errors.Is(e.appErr, ErrInvalidObjectID):
		return http.StatusBadRequest
	case errors.Is(e.appErr, ErrValidation):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
