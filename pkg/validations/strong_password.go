package validations

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// StrongPassword Custom validation function for strong passwords
func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Check for minimum length
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		// You can adjust the set of allowed special characters as needed.
		case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~", char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
