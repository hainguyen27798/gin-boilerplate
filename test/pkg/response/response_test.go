package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestResponsePackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Response Package Suite")
}

var _ = Describe("Response Handlers", func() {
	var (
		c *gin.Context
		w *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
	})

	Describe("OkResponse", func() {
		It("should return a successful response with data", func() {
			data := map[string]string{"key": "value"}
			message := "Success"

			response.OkResponse(c, message, data)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(message))
			Expect(w.Body.String()).To(ContainSubstring("\"key\":\"value\""))
		})
	})

	Describe("CreatedResponse", func() {
		It("should return a created response with data", func() {
			data := map[string]string{"id": "123"}
			message := "Resource created"

			response.CreatedResponse(c, message, data)

			Expect(w.Code).To(Equal(http.StatusCreated))
			Expect(w.Body.String()).To(ContainSubstring(message))
			Expect(w.Body.String()).To(ContainSubstring("\"id\":\"123\""))
		})
	})

	Describe("ValidateErrorResponse", func() {
		Context("with validation errors", func() {
			It("should format and return validation errors properly", func() {
				// Create mock validation errors
				var validationErrors validator.ValidationErrors
				validationErrors = append(validationErrors, MockValidationErrors{})

				response.ValidateErrorResponse(c, validationErrors)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(ContainSubstring("validation error"))
				Expect(w.Body.String()).To(ContainSubstring("Field username is required"))
			})
		})

		Context("with regular errors", func() {
			It("should return a validation error response", func() {
				err := errors.New("some error")
				response.ValidateErrorResponse(c, err)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(ContainSubstring("validation error"))
			})
		})

		Context("with struct validation errors", func() {
			It("should format all validation errors correctly", func() {
				validate := validator.New()
				testStruct := MockStruct{
					Name:    "",
					Email:   "invalid-email",
					Age:     15,
					Website: "not-a-url",
				}

				err := validate.Struct(testStruct)
				response.ValidateErrorResponse(c, err)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(ContainSubstring("Field Name is required"))
				Expect(w.Body.String()).To(ContainSubstring("Fields Email is invalid due to email"))
				Expect(w.Body.String()).To(ContainSubstring("Fields Age is invalid due to min"))
				Expect(w.Body.String()).To(ContainSubstring("Fields Website is invalid due to url"))
			})
		})
	})

	Describe("ErrorResponse", func() {
		errorTestCases := []struct {
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

		for _, tc := range errorTestCases {
			// Use local variables to avoid closure issues
			testCase := tc

			Context("with "+testCase.name, func() {
				It("should return the correct status code and error messages", func() {
					// Create a fresh context for each test
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					err := response.NewError(testCase.appErr, testCase.serviceErr)
					response.ErrorResponse(c, err)

					Expect(w.Code).To(Equal(testCase.expectedStatus))

					for _, expectedStr := range testCase.expectedBody {
						Expect(w.Body.String()).To(ContainSubstring(expectedStr))
					}
				})
			})
		}
	})

	Describe("MessageResponse", func() {
		messageTestCases := []struct {
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

		for _, tc := range messageTestCases {
			// Use local variables to avoid closure issues
			testCase := tc

			Context("with "+testCase.name, func() {
				It("should return the correct error message", func() {
					// Create a fresh context for each test
					w := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(w)

					response.ErrorResponse(c, response.NewError(testCase.appErr, nil))

					Expect(w.Code).To(Equal(testCase.expectedStatus))
					Expect(w.Body.String()).To(ContainSubstring(testCase.expectedMsg))
				})
			})
		}
	})
})

// Keep the mock types as they were
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

type MockStruct struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"email"`
	Age     int    `json:"age" validate:"min=18"`
	Website string `json:"website" validate:"url"`
}
