package users

import (
	"context"
	"crypto/rand"

	"github.com/hainguyen27798/gin-boilerplate/pkg/response"

	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// UserService defines the interface for user-related operations.
type UserService interface {
	CreateUser(ctx context.Context, user *CreateUserDto) (*UserDto, *response.Error)
	GetUserByEmail(ctx context.Context, email string) (*UserDto, *response.Error)
	GetUserByID(ctx context.Context, id string) (*UserDto, *response.Error)
	UpdateUser(ctx context.Context, id string, user *UpdateUserDto) (*UserDto, *response.Error)
	DeleteUser(ctx context.Context, id string) *response.Error
}

// userServiceImpl is the concrete implementation of UserService
type userServiceImpl struct {
	repo UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

// CreateUser creates a new user with validation and additional processing
func (s *userServiceImpl) CreateUser(
	ctx context.Context,
	user *CreateUserDto,
) (*UserDto, *response.Error) {
	// Hash password before storing
	passwordHashed, err := helpers.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := &UserModel{
		Email:            user.Email,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Password:         passwordHashed,
		Image:            user.Image,
		VerificationCode: generateVerificationCode(),
		Verified:         false,
	}

	// Create user in repository
	userCreated, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return userCreated.ToDto(), err
}

// GetUserByEmail retrieves a user by their email address
func (s *userServiceImpl) GetUserByEmail(
	ctx context.Context,
	email string,
) (*UserDto, *response.Error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user.ToDto(), nil
}

// GetUserByID retrieves a user by their ID
func (s *userServiceImpl) GetUserByID(ctx context.Context, id string) (*UserDto, *response.Error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.ToDto(), nil
}

// UpdateUser updates an existing user
func (s *userServiceImpl) UpdateUser(
	ctx context.Context, id string,
	user *UpdateUserDto,
) (*UserDto, *response.Error) {
	// Find the user by ID
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update user in repository
	newDoc := helpers.MustValue(common.ToBson(user))

	userUpdated, err := s.repo.Update(ctx, id, bson.D{{Key: "$set", Value: newDoc}})
	if err != nil {
		return nil, err
	}

	return userUpdated.ToDto(), nil
}

// DeleteUser deletes a user by their ID
func (s *userServiceImpl) DeleteUser(ctx context.Context, id string) *response.Error {
	return s.repo.Delete(ctx, id)
}

// Helper function to generate verification code
func generateVerificationCode() string {
	otpChars := "1234567890"
	length := 6
	buffer := make([]byte, length)
	helpers.MustValue(rand.Read(buffer))

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}
	return string(buffer)
}
