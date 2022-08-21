package repositories

import (
	"context"
	"veterinary-employee/db"
	"veterinary-employee/models"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AvatarRepository struct {
	Data *db.Data
}

var avatarCollection = settings.InitializeMongoDB().Collections.Avatar

func (r *AvatarRepository) Get(ctx context.Context, filter interface{}) (models.Avatar, error) {
	collection := r.Data.DB.Collection(avatarCollection)
	singleResult := collection.FindOne(ctx, filter)

	var avatar models.Avatar

	if err := singleResult.Decode(&avatar); err != nil {
		return avatar, err
	}

	return avatar, nil
}

func (r *AvatarRepository) Create(ctx context.Context, avatar models.Avatar) (models.Avatar, error) {
	collection := r.Data.DB.Collection(avatarCollection)

	if _, err := collection.InsertOne(ctx, avatar); err != nil {
		return avatar, err
	}

	return avatar, nil
}

func (r *AvatarRepository) Update(
	ctx context.Context,
	filter interface{},
	document interface{},
) (models.Avatar, error) {
	collection := r.Data.DB.Collection(avatarCollection)

	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	singleResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": document}, &updateOptions)

	var avatar models.Avatar

	if err := singleResult.Decode(&avatar); err != nil {
		return avatar, err
	}

	return avatar, nil
}

func (r *AvatarRepository) Delete(ctx context.Context, filter interface{}) error {
	collection := r.Data.DB.Collection(avatarCollection)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (r *AvatarRepository) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	collection := r.Data.DB.Collection(avatarCollection)
	return collection.CountDocuments(ctx, filter)
}
