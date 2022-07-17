package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatar struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EmployeeId primitive.ObjectID `json:"employee_id" bson:"employee_id,omitempty"`
	Path       string             `json:"path" bson:"path,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

func (a *Avatar) ValidatePartial() error {
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Avatar) Validate() error {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	return nil
}
