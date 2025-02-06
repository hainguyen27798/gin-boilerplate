package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnCode(t *testing.T) {
	t.Run("should create ServerCode with given code", func(t *testing.T) {
		code := ReturnCode(CodeSuccess)
		assert.Equal(t, CodeSuccess, code.Code())
	})
}

func TestServerCode_InValid(t *testing.T) {
	t.Run("should return false for success codes", func(t *testing.T) {
		successCodes := []int{
			CodeSuccess,
			LoginSuccess,
			LogoutSuccess,
			CreatedSuccess,
		}

		for _, code := range successCodes {
			sc := ReturnCode(code)
			assert.False(t, sc.InValid(), "Expected code %d to be valid", code)
		}
	})

	t.Run("should return true for error codes", func(t *testing.T) {
		errorCodes := []int{
			ErrBadRequest,
			ErrCodeParamInvalid,
			ErrCreateFailed,
			ErrInvalidOTP,
			ErrUnauthorized,
			ErrNotFound,
			ErrInternalError,
		}

		for _, code := range errorCodes {
			sc := ReturnCode(code)
			assert.True(t, sc.InValid(), "Expected code %d to be invalid", code)
		}
	})
}

func TestCodeMsg(t *testing.T) {
	t.Run("should have matching messages for all defined codes", func(t *testing.T) {
		codes := []int{
			CodeSuccess,
			LoginSuccess,
			LogoutSuccess,
			CreatedSuccess,
			ErrBadRequest,
			ErrCodeParamInvalid,
			ErrCreateFailed,
			ErrInvalidOTP,
			ErrSendEmailFailed,
			ErrCodeUserHasExists,
			ErrCodeLoginFailed,
			ErrUnauthorized,
			ErrInvalidToken,
			ErrExpiredToken,
			ErrStolenToken,
			ErrNotFound,
			ErrCodeUserNotExists,
			ErrInternalError,
			ErrJWTInternalError,
		}

		for _, code := range codes {
			msg, exists := CodeMsg[code]
			assert.True(t, exists, "Expected message to exist for code %d", code)
			assert.NotEmpty(t, msg, "Expected non-empty message for code %d", code)
		}
	})

	t.Run("should have unique messages for each code", func(t *testing.T) {
		seen := make(map[string]int)
		for code, msg := range CodeMsg {
			if existing, exists := seen[msg]; exists {
				t.Errorf("Duplicate message '%s' found for codes %d and %d", msg, existing, code)
			}
			seen[msg] = code
		}
	})
}
