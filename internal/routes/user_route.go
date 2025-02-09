package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/gin-boilerplate/internal/module/users"
)

// RegisterUserRoutes sets up the routes for user-related operations.
func RegisterUserRoutes(router *gin.Engine, userController *users.UserController) {
	// Group user-related routes
	userRoutes := router.Group("v1/users")
	{
		userRoutes.POST("", userController.CreateUser)       // Create a new user
		userRoutes.GET("/:id", userController.GetUserByID)   // Get a user by ID
		userRoutes.GET("", userController.GetUserByEmail)    // Get a user by email
		userRoutes.PUT("/:id", userController.UpdateUser)    // Update a user
		userRoutes.DELETE("/:id", userController.DeleteUser) // Delete a user
	}
}
