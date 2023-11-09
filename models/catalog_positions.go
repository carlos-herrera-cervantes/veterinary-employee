package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatalogPosition struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"required"`
	Active      bool               `json:"active" bson:"active"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type PartialCatalogPosition struct {
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Active      *bool     `json:"active" bson:"active,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func (cp *PartialCatalogPosition) Validate() error {
	cp.UpdatedAt = time.Now().UTC()
	return nil
}

func (cp *CatalogPosition) Validate() error {
	cp.Id = primitive.NewObjectID()
	cp.Active = true
	cp.CreatedAt = time.Now()
	cp.UpdatedAt = time.Now()

	var structValidator = validator.New()
	structError := structValidator.Struct(cp)

	if structError != nil {
		return structError
	}

	return nil
}
