package repositories

import (
	"context"
	"veterinary-employee/db"
	"veterinary-employee/models"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AddressRepository struct {
	Data *db.Data
}

var addressCollection = settings.InitializeMongoDB().Collections.Address

func (r *AddressRepository) Get(ctx context.Context, filter interface{}) (models.Address, error) {
	collection := r.Data.DB.Collection(addressCollection)
	singleResult := collection.FindOne(ctx, filter)

	var address models.Address

	if err := singleResult.Decode(&address); err != nil {
		return address, err
	}

	return address, nil
}

func (r *AddressRepository) Create(ctx context.Context, address models.Address) (models.Address, error) {
	collection := r.Data.DB.Collection(addressCollection)

	if _, err := collection.InsertOne(ctx, address); err != nil {
		return address, err
	}

	return address, nil
}

func (r *AddressRepository) Update(
	ctx context.Context,
	filter interface{},
	document interface{},
) (models.Address, error) {
	collection := r.Data.DB.Collection(addressCollection)

	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	singleResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": document}, &updateOptions)

	var address models.Address

	if err := singleResult.Decode(&address); err != nil {
		return address, err
	}

	return address, nil
}
