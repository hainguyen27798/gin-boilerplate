//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/gin-boilerplate/internal/module/users"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// InitializeUserModule sets up the UserController with its dependencies.
func InitializeUserModule(db *mongo.Database) *users.UserController {
	wire.Build(
		users.NewUserRepository,
		users.NewUserService,
		users.NewUserController,
	)
	return &users.UserController{}
}
