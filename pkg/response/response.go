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

// MessageResponse is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status code determined by the provided code parameter. The response includes
// the provided code and the corresponding message from the CodeMsg map.
func MessageResponse(c *gin.Context, code int) {
	c.JSON(GetHTTPCode(code), TResponse{
		Code:    code,
		Message: CodeMsg[code],
	})
}

// NotFoundException is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status of http.StatusNotFound. The response includes the provided code and
// the corresponding message from the CodeMsg map.
// This function also calls c.Abort() to stop the current request processing.
func NotFoundException(c *gin.Context, code int) {
	c.JSON(http.StatusNotFound, TResponse{
		Code:    code,
		Message: CodeMsg[code],
	})
	c.Abort()
}

// OkResponse is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status of http.StatusOK. The response includes the provided code and
// the corresponding message from the CodeMsg map, as well as the provided data.
func OkResponse[T any](c *gin.Context, code int, data T) {
	c.JSON(http.StatusOK, TDataResponse{
		TResponse: TResponse{
			Code:    code,
			Message: CodeMsg[code],
		},
		Data: data,
	})
}

// CreatedResponse is a helper function that writes a JSON response to the provided gin.Context
// with an HTTP status of http.StatusCreated. The response includes the provided code and
// the corresponding message from the CodeMsg map, as well as the provided data.
func CreatedResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusCreated, TDataResponse{
		TResponse: TResponse{
			Code:    code,
			Message: CodeMsg[code],
		},
		Data: data,
	})
}

// ValidateErrorResponse is a helper function that writes a JSON response to the
// provided gin.Context with an HTTP status of http.StatusBadRequest. The
// response includes the error code ErrCodeParamInvalid and the validation
// errors returned by the validationErrorsToJSON function.
func ValidateErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, TErrResponse{
		TResponse: TResponse{
			Code:    ErrCodeParamInvalid,
			Message: CodeMsg[ErrCodeParamInvalid],
		},
		Errors: validationErrorsToJSON(err),
	})
}

// GetHTTPCode maps a response code to the corresponding HTTP status code.
// It handles the following cases:
// - Codes in the range [20000, 20100] map to http.StatusOK
// - Codes in the range [20100, 20200] map to http.StatusOK
// - Codes in the range [40000, 50000] map to http.StatusBadRequest
// - All other codes map to http.StatusInternalServerError
func GetHTTPCode(code int) int {
	switch {
	case code >= 20000 && code < 20100:
		return http.StatusOK
	case code >= 20100 && code < 20200:
		return http.StatusOK
	case code >= 40000 && code < 50000:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
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
