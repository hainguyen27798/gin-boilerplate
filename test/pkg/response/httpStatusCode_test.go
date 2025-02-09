package response

import (
	response2 "github.com/hainguyen27798/gin-boilerplate/pkg/response"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnCode(t *testing.T) {
	t.Run("should create ServerCode with given code", func(t *testing.T) {
		code := response2.ReturnCode(response2.CodeSuccess)
		assert.Equal(t, response2.CodeSuccess, code.Code())
	})
}

func TestServerCode_InValid(t *testing.T) {
	t.Run("should return false for success codes", func(t *testing.T) {
		successCodes := []int{
			response2.CodeSuccess,
			response2.LoginSuccess,
			response2.LogoutSuccess,
			response2.CreatedSuccess,
		}

		for _, code := range successCodes {
			sc := response2.ReturnCode(code)
			assert.False(t, sc.InValid(), "Expected code %d to be valid", code)
		}
	})

	t.Run("should return true for error codes", func(t *testing.T) {
		errorCodes := []int{
			response2.ErrBadRequest,
			response2.ErrCodeParamInvalid,
			response2.ErrCreateFailed,
			response2.ErrInvalidOTP,
			response2.ErrUnauthorized,
			response2.ErrNotFound,
			response2.ErrInternalError,
		}

		for _, code := range errorCodes {
			sc := response2.ReturnCode(code)
			assert.True(t, sc.InValid(), "Expected code %d to be invalid", code)
		}
	})
}

func TestCodeMsg(t *testing.T) {
	t.Run("should have matching messages for all defined codes", func(t *testing.T) {
		codes := []int{
			response2.CodeSuccess,
			response2.LoginSuccess,
			response2.LogoutSuccess,
			response2.CreatedSuccess,
			response2.ErrBadRequest,
			response2.ErrCodeParamInvalid,
			response2.ErrCreateFailed,
			response2.ErrInvalidOTP,
			response2.ErrSendEmailFailed,
			response2.ErrCodeUserHasExists,
			response2.ErrCodeLoginFailed,
			response2.ErrUnauthorized,
			response2.ErrInvalidToken,
			response2.ErrExpiredToken,
			response2.ErrStolenToken,
			response2.ErrNotFound,
			response2.ErrCodeUserNotExists,
			response2.ErrInternalError,
			response2.ErrJWTInternalError,
		}

		for _, code := range codes {
			msg, exists := response2.CodeMsg[code]
			assert.True(t, exists, "Expected message to exist for code %d", code)
			assert.NotEmpty(t, msg, "Expected non-empty message for code %d", code)
		}
	})

	t.Run("should have unique messages for each code", func(t *testing.T) {
		seen := make(map[string]int)
		for code, msg := range response2.CodeMsg {
			if existing, exists := seen[msg]; exists {
				t.Errorf("Duplicate message '%s' found for codes %d and %d", msg, existing, code)
			}
			seen[msg] = code
		}
	})
}
