package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name" validate:"required"`
	Active    bool               `json:"active" bson:"active"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type PartialRole struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	Active    *bool     `json:"active" bson:"active,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (r *PartialRole) ValidatePartial() error {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Role) Validate() error {
	r.Id = primitive.NewObjectID()
	r.Active = true
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	var structValidator = validator.New()
	structError := structValidator.Struct(r)

	if structError != nil {
		return structError
	}

	return nil
}
