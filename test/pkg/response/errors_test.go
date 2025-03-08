package response

import (
	"errors"
	"net/http"
	"testing"

	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestErrorsPackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errors Package Suite")
}

var _ = Describe("Error Handling", func() {
	Describe("NewError", func() {
		It("should create a new error with application and service errors", func() {
			appErr := errors.New("application error")
			serviceErr := errors.New("service error")

			err := response.NewError(appErr, serviceErr)

			Expect(err).NotTo(BeNil())
			Expect(err.AppErr()).To(Equal(appErr.Error()))
			Expect(err.ServiceErr()).To(Equal(serviceErr.Error()))
		})
	})

	Describe("Error.Error()", func() {
		It("should return a string containing both error messages", func() {
			appErr := errors.New("application error")
			serviceErr := errors.New("service error")

			err := response.NewError(appErr, serviceErr)

			Expect(err.Error()).To(ContainSubstring("application error"))
			Expect(err.Error()).To(ContainSubstring("service error"))
		})
	})

	Describe("Error.Code()", func() {
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
			// Use local variable to avoid closure issues
			appErr := tc.appErr
			expectedCode := tc.expected
			name := tc.name

			It("should return the correct status code for "+name, func() {
				err := response.NewError(appErr, errors.New("service error"))
				Expect(err.Code()).To(Equal(expectedCode))
			})
		}
	})

	Describe("Error.AppErr()", func() {
		It("should return the application error message", func() {
			appErr := response.ErrNotFound
			serviceErr := errors.New("service error")

			err := response.NewError(appErr, serviceErr)

			Expect(err.AppErr()).To(Equal(appErr.Error()))
		})
	})

	Describe("Error.ServiceErr()", func() {
		It("should return the service error message", func() {
			appErr := response.ErrBadRequest
			serviceErr := errors.New("service error")

			err := response.NewError(appErr, serviceErr)

			Expect(err.ServiceErr()).To(Equal(serviceErr.Error()))
		})
	})
})
