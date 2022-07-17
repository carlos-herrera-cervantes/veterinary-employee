package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	EmployeeId   primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Municipality string             `json:"municipality" bson:"municipality"`
	PostalCode   string             `json:"postal_code" bson:"postal_code"`
	Street       string             `json:"street" bson:"street"`
	Colony       string             `json:"colony" bson:"colony"`
	Number       string             `json:"number" bson:"number"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type PartialAddress struct {
	Municipality string    `json:"municipality" bson:"municipality,omitempty"`
	PostalCode   string    `json:"postal_code" bson:"postal_code,omitempty"`
	Street       string    `json:"street" bson:"street,omitempty"`
	Colony       string    `json:"colony" bson:"colony,omitempty"`
	Number       string    `json:"number" bson:"number,omitempty"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

func (a *PartialAddress) ValidatePartial() error {
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Address) Validate() error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}
