package repositories

import (
	"context"
	"veterinary-employee/db"
	"veterinary-employee/models"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileRepository struct {
	Data *db.Data
}

var profileCollection = settings.InitializeMongoDB().Collections.Profile

func (r *ProfileRepository) GetAll(ctx context.Context, page, pageSize int64) ([]models.Profile, error) {
	collection := r.Data.DB.Collection(profileCollection)

	skip := page * pageSize
	findOptions := options.FindOptions{Limit: &pageSize, Skip: &skip}
	cursor, err := collection.Find(ctx, bson.D{}, &findOptions)

	var profiles []models.Profile

	if err != nil {
		return profiles, err
	}

	if err := cursor.All(ctx, &profiles); err != nil {
		return profiles, err
	}

	return profiles, nil
}

func (r *ProfileRepository) Get(ctx context.Context, filter interface{}) (models.Profile, error) {
	collection := r.Data.DB.Collection(profileCollection)
	singleResult := collection.FindOne(ctx, filter)

	var profile models.Profile

	if err := singleResult.Decode(&profile); err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *ProfileRepository) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	collection := r.Data.DB.Collection(profileCollection)
	total, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *ProfileRepository) Update(
	ctx context.Context,
	filter interface{},
	document interface{},
) (models.Profile, error) {
	collection := r.Data.DB.Collection(profileCollection)

	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	singleResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": document}, &updateOptions)

	var profile models.Profile

	if err := singleResult.Decode(&profile); err != nil {
		return profile, err
	}

	return profile, nil
}
