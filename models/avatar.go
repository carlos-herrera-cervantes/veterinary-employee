package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatar struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	EmployeeId primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Path       string             `json:"path" bson:"path"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
