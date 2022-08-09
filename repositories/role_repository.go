package repositories

import (
	"context"
	"veterinary-employee/db"
	"veterinary-employee/models"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RoleRepository struct {
	Data *db.Data
}

var roleCollection = settings.InitializeMongoDB().Collections.Role

func (r *RoleRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	collection := r.Data.DB.Collection(roleCollection)
	cursor, err := collection.Find(ctx, bson.D{})

	var roles []models.Role

	if err != nil {
		return roles, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return roles, err
	}

	return roles, nil
}

func (r *RoleRepository) Get(ctx context.Context, filter interface{}) (models.Role, error) {
	collection := r.Data.DB.Collection(roleCollection)
	singleResult := collection.FindOne(ctx, filter)

	var role models.Role

	if err := singleResult.Decode(&role); err != nil {
		return role, err
	}

	return role, nil
}

func (r *RoleRepository) Create(ctx context.Context, role models.Role) (models.Role, error) {
	collection := r.Data.DB.Collection(roleCollection)

	if _, err := collection.InsertOne(ctx, role); err != nil {
		return role, err
	}

	return role, nil
}

func (r *RoleRepository) Update(
	ctx context.Context,
	filter interface{},
	document interface{},
) (models.Role, error) {
	collection := r.Data.DB.Collection(roleCollection)

	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	singleResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": document}, &updateOptions)

	var role models.Role

	if err := singleResult.Decode(&role); err != nil {
		return role, err
	}

	return role, nil
}
