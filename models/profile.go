package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	EmployeeId  primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Email       string             `json:"email" bson:"email"`
	Name        string             `json:"name" bson:"name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Gender      string             `json:"gender" bson:"gender"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Birthday    string             `json:"birthday" bson:"birthday"`
	Roles       []string           `json:"roles" bson:"roles"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
