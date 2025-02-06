package common

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// BaseModel provides common fields and methods for all models in the application
type BaseModel struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

// BeforeCreate sets the creation timestamp
func (m *BaseModel) BeforeCreate() {
	if m.ID.IsZero() {
		m.ID = bson.NewObjectID()
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().UTC()
	}
	m.UpdatedAt = time.Now().UTC()
}

// BeforeUpdate updates the modification timestamp
func (m *BaseModel) BeforeUpdate() {
	m.UpdatedAt = time.Now().UTC()
}
