package users

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
)

// CreateUserDto is used for creating a new user.
type CreateUserDto struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Password  string `json:"password" validate:"required,strongPassword"`
	Image     string `json:"image" validate:"omitempty,url"`
}

// Validate validates the CreateUserDto.
func (dto *CreateUserDto) Validate() error {
	return common.ValidateStruct(dto)
}

// UpdateUserDto is used for updating an existing user.
type UpdateUserDto struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Image     string `json:"image,omitempty" bson:"image,omitempty"`
}

// Validate validates the UpdateUserDto.
func (dto *UpdateUserDto) Validate() error {
	return common.ValidateStruct(dto)
}

// UserDto is used for retrieving user information, excluding the password.
type UserDto struct {
	common.BaseDto `json:",inline"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Image          string `json:"image"`
	Verified       bool   `json:"verified"`
}
