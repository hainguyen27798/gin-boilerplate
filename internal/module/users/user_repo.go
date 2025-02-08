package users

import (
	"context"
	"time"

	"github.com/hainguyen27798/gin-boilerplate/pkg/helpers"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UserRepository defines the interface for user repository operations.
type UserRepository interface {
	Create(ctx context.Context, user *UserModel) (*UserModel, error)
	FindByEmail(ctx context.Context, email string) (*UserModel, error)
	FindByID(ctx context.Context, id string) (*UserModel, error)
	Update(ctx context.Context, id string, payload bson.D) (*UserModel, error)
	Delete(ctx context.Context, id string) error
}

// userRepositoryImpl is a concrete implementation of UserRepository
type userRepositoryImpl struct {
	model *mongo.Collection
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		model: db.Collection(UserModel{}.CollectionName()),
	}
}

// Create inserts a new user into the database. It first calls the BeforeCreate method on the user
// model, then inserts the user into the database. If the insert operation is successful, it returns
// the created user model. If there is an error, it returns the error.
func (r *userRepositoryImpl) Create(ctx context.Context, user *UserModel) (*UserModel, error) {
	user.BeforeCreate()
	if _, err := r.model.InsertOne(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail retrieves a user by their email address
func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*UserModel, error) {
	var user UserModel
	err := r.model.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by their ID
func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*UserModel, error) {
	var user UserModel
	_id := helpers.MustValue(bson.ObjectIDFromHex(id))
	err := r.model.FindOne(ctx, bson.M{"_id": _id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates an existing user in the database
func (r *userRepositoryImpl) Update(
	ctx context.Context,
	id string, payload bson.D,
) (*UserModel, error) {
	_id := helpers.MustValue(bson.ObjectIDFromHex(id))

	// Append "updated_at" field to the payload.
	payload = append(payload, bson.E{
		Key: "$set",
		Value: bson.D{
			{Key: "updated_at", Value: time.Now()},
		},
	})

	var userUpdated UserModel
	err := r.model.FindOneAndUpdate(
		ctx,
		bson.M{"_id": _id},
		payload,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&userUpdated)
	if err != nil {
		return nil, err
	}

	return &userUpdated, err
}

// Delete removes a user from the database by their ID
func (r *userRepositoryImpl) Delete(ctx context.Context, id string) error {
	_id := helpers.MustValue(bson.ObjectIDFromHex(id))
	_, err := r.model.DeleteOne(ctx, bson.M{"_id": _id})
	return err
}
