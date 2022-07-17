package models

import (
	"time"

	"veterinary-employee/enums/gender"
	"veterinary-employee/enums/role"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	EmployeeId  primitive.ObjectID `json:"employee_id" bson:"employee_id"`
	Email       string             `json:"email" bson:"email"`
	Name        string             `json:"name" bson:"name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Gender      string             `json:"gender" bson:"gender" validate:"validateGenders"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	Birthday    string             `json:"birthday" bson:"birthday"`
	Roles       []string           `json:"roles" bson:"roles" validate:"validateRoles"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type PartialProfile struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	LastName  string    `json:"last_name" bson:"last_name,omitempty"`
	Gender    string    `json:"gender" bson:"gender,omitempty"`
	Birthday  string    `json:"birthday" bson:"birthday,omitempty"`
	Roles     []string  `json:"roles" bson:"roles,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

func validateGenderAttribute(fl validator.FieldLevel) bool {
	genderField := fl.Field().String()
	validGenders := map[string]bool{
		gender.Female:       true,
		gender.Male:         true,
		gender.NotSpecified: true,
	}

	return validGenders[genderField]
}

func validateRoleAttribute(fl validator.FieldLevel) bool {
	roleField := fl.Field().String()
	validRoles := map[string]bool{
		role.Employee:   true,
		role.SuperAdmin: true,
	}

	return validRoles[roleField]
}

func (p *PartialProfile) ValidatePartial() error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Profile) Validate() error {
	var structValidator = validator.New()

	structValidator.RegisterValidation("validateGenders", validateGenderAttribute)
	structValidator.RegisterValidation("validateRoles", validateRoleAttribute)
	structError := structValidator.Struct(p)

	if structError != nil {
		return structError
	}

	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return nil
}
