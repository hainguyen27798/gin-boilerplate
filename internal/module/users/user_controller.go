package users

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/pkg/common"
	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"github.com/hainguyen27798/gin-boilerplate/pkg/response"
)

// UserController handles HTTP requests related to user operations.
type UserController struct {
	userService UserService
}

// NewUserController creates a new instance of UserController.
func NewUserController(userService UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser handles the creation of a new user.
func (c *UserController) CreateUser(ctx *gin.Context) {
	var dto CreateUserDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.ValidateErrorResponse(ctx, err)
		return
	}

	if err := dto.Validate(); err != nil {
		response.ValidateErrorResponse(ctx, err)
		return
	}

	newUser, err := c.userService.CreateUser(ctx, &dto)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.CreatedResponse(ctx, "Created user successfully", newUser)
}

// GetUserByID handles the retrieval of a user by ID.
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if ok := common.IsValidObjectID(id); !ok {
		response.ErrorResponse(ctx, response.NewError(response.ErrInvalidObjectID, nil))
		return
	}

	user, err := c.userService.GetUserByID(ctx, id)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.OkResponse(ctx, "Found a user", user)
}

// UpdateUser handles the update of a user.
func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if ok := common.IsValidObjectID(id); !ok {
		response.ErrorResponse(ctx, response.NewError(response.ErrInvalidObjectID, nil))
		return
	}

	var dto UpdateUserDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.ValidateErrorResponse(ctx, err)
		return
	}

	if err := dto.Validate(); err != nil {
		response.ValidateErrorResponse(ctx, err)
		return
	}

	userDto, err := c.userService.UpdateUser(ctx, id, &dto)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.OkResponse(ctx, "Updated user successfully", userDto)
}

// DeleteUser handles the deletion of a user.
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.ErrorResponse(ctx, response.NewError(response.ErrInvalidObjectID, nil))
		return
	}

	if err := c.userService.DeleteUser(ctx, id); err != nil {
		helpers.Must(err)
		response.ErrorResponse(ctx, err)
		return
	}

	response.OkResponse(ctx, "Deleted user successfully", nil)
}

// GetUserByEmail handles the retrieval of a user by email.
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		response.ErrorResponse(ctx, response.NewError(response.ErrInvalidObjectID, nil))
		return
	}

	user, err := c.userService.GetUserByEmail(ctx, email)
	if err != nil {
		response.ErrorResponse(ctx, err)
		return
	}

	response.OkResponse(ctx, "Found a user", user)
}
