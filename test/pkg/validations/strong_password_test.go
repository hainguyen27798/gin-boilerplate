package validations

import (
	validations2 "github.com/hainguyen27798/gin-boilerplate/pkg/validations"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Password string `validate:"strong_password"`
}

func TestStrongPassword(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("strong_password", validations2.StrongPassword)
	if err != nil {
		return
	}

	t.Run("should accept valid passwords", func(t *testing.T) {
		validPasswords := []string{
			"Ab@12345",
			"SecurePass123!",
			"ValidP@ssw0rd",
		}

		for _, password := range validPasswords {
			test := TestStruct{Password: password}
			err := validate.Struct(test)
			assert.NoError(t, err, "Password should be valid: %s", password)
		}
	})

	t.Run("should reject invalid passwords", func(t *testing.T) {
		invalidPasswords := []string{
			"",
			"short",
			"1234567",
			"12345678",
			"!@#$%^&*",
			"abcdefg",
			"       ",
			"\t\n\r",
		}

		for _, password := range invalidPasswords {
			test := TestStruct{Password: password}
			err := validate.Struct(test)
			assert.Error(t, err, "Password should be invalid: %s", password)
		}
	})

	t.Run("should handle edge cases", func(t *testing.T) {
		edgeCases := []struct {
			password string
			valid    bool
		}{
			{"12345678", false},                         // Numbers only - invalid
			{"123456789012345678901234567890", false},   // Numbers only - invalid
			{"        ", false},                         // Spaces only - invalid
			{"!@#$%^&*", false},                         // Special chars only - invalid
			{"ñÑáéíóúÁÉÍÓÚ12345678", false},             // Unicode without required char mix
			{"\x00\x00\x00\x00\x00\x00\x00\x00", false}, // Null bytes
			{"Str0ng!Pass", true},                       // Valid mix of characters
			{"Complex1@Pwd", true},                      // Valid mix
		}

		for _, tc := range edgeCases {
			test := TestStruct{Password: tc.password}
			err := validate.Struct(test)
			if tc.valid {
				assert.NoError(t, err, "Password should be valid: %s", tc.password)
			} else {
				assert.Error(t, err, "Password should be invalid: %s", tc.password)
			}
		}
	})
}
