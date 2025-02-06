package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/gin-boilerplate/global"
	"github.com/hainguyen27798/gin-boilerplate/pkg/validations"
)

// RegisterValidations Register custom validation
func RegisterValidations() {
	v := validator.New()

	// Register custom strong password validation
	if err := v.RegisterValidation("strongPassword", validations.StrongPassword); err != nil {
		panic(err)
	}

	global.Validator = v
}
