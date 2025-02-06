package users

import (
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
)

// UserModel represents the data structure for a user in the system.
type UserModel struct {
	common.BaseModel `bson:",inline"`
	Email            string `bson:"email,omitempty" json:"email,omitempty"`
	FirstName        string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName         string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Password         string `bson:"password,omitempty" json:"-"`
	Image            string `bson:"image" json:"image"`
	Verified         bool   `bson:"verified" json:"verified"`
	VerificationCode string `bson:"verification_code" json:"-"`
}

// CollectionName returns the name of the MongoDB collection for this model.
func (UserModel) CollectionName() string {
	return "users"
}

// ToDto returns the name of the MongoDB collection for this model.
func (user UserModel) ToDto() *UserDto {
	return &UserDto{
		ID:        user.ID.Hex(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Image:     user.Image,
		Verified:  user.Verified,
	}
}
