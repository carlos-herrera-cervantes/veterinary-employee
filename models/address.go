package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	EmployeeId   primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Municipality string             `json:"municipality" bson:"municipality"`
	PostalCode   string             `json:"postal_code" bson:"postal_code"`
	Street       string             `json:"street" bson:"street"`
	Colony       string             `json:"colony" bson:"colony"`
	Number       string             `json:"number" bson:"number"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
