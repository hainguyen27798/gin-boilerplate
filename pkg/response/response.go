package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// TResponse is a struct that represents a basic response from the API. It contains
// a code field to indicate the response status and a message field to provide
// additional context.
type TResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TDataResponse is a response type that includes a data field. It embeds the TResponse
// struct and adds a Data field to hold the response data.
type TDataResponse struct {
	TResponse
	Data interface{} `json:"data"`
}

// TErrResponse is a response type that includes an error field. It embeds the TResponse
// struct and adds an Errors field to hold any errors that occurred.
type TErrResponse struct {
	TResponse
	Errors interface{} `json:"errors"`
}

// OkResponse is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status of http.StatusOK. The response includes the provided message and
// the provided data.
func OkResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, TDataResponse{
		TResponse: TResponse{
			Code:    http.StatusOK,
			Message: msg,
		},
		Data: data,
	})
}

// CreatedResponse is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status of http.StatusCreated. The response includes the provided message and
// the provided data.
func CreatedResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusCreated, TDataResponse{
		TResponse: TResponse{
			Code:    http.StatusCreated,
			Message: msg,
		},
		Data: data,
	})
}

// ErrorResponse writes a JSON response to the provided gin.Context with an HTTP status
// corresponding to the error code. The response includes the error message.
func ErrorResponse(c *gin.Context, err *Error) {
	c.JSON(err.Code(), TErrResponse{
		TResponse: TResponse{
			Code:    err.Code(),
			Message: err.AppErr(),
		},
		Errors: err.ServiceErr(),
	})
}

// ValidateErrorResponse writes a JSON response to the provided gin.Context with an HTTP status
// corresponding to the error code. The response includes the error message and a list of
// validation error messages, if any.
func ValidateErrorResponse(c *gin.Context, err error) {
	errValidation := NewError(ErrValidation, err)
	c.JSON(errValidation.Code(), TErrResponse{
		TResponse: TResponse{
			Code:    errValidation.Code(),
			Message: errValidation.AppErr(),
		},
		Errors: validationErrorsToJSON(err),
	})
}

// validationErrorsToJSON is a helper function that converts a validator.ValidationErrors
// error to a slice of strings containing the error messages. It handles both "required"
// and other validation error types.
func validationErrorsToJSON(err error) []string {
	var res []string

	// Check if it's a validation error
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, fieldErr := range validationErrs {
			// Collect error messages into a list
			if fieldErr.Tag() == "required" {
				res = append(res, fmt.Sprintf("Field %s is required", fieldErr.Field()))
			} else {
				res = append(
					res,
					fmt.Sprintf("Fields %s is invalid due to %s", fieldErr.Field(), fieldErr.Tag()),
				)
			}
		}
	}

	return res
}
