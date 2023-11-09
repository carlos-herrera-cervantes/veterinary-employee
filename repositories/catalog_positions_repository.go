package repositories

import (
	"context"

	"veterinary-employee/db"
	"veterinary-employee/models"
	"veterinary-employee/settings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatalogPositionsRepository struct {
	Data *db.Data
}

var catalogPositionsCollection = settings.InitializeMongoDB().Collections.CatalogPositions

func (r CatalogPositionsRepository) GetAll(ctx context.Context, page, pageSize int64) ([]models.CatalogPosition, error) {
	collection := r.Data.DB.Collection(catalogPositionsCollection)
	skip := page * pageSize
	findOptions := options.FindOptions{Limit: &pageSize, Skip: &skip}
	cursor, err := collection.Find(ctx, bson.D{}, &findOptions)

	var catalogPositions []models.CatalogPosition

	if err != nil {
		return catalogPositions, err
	}

	if err := cursor.All(ctx, &catalogPositions); err != nil {
		return catalogPositions, err
	}

	return catalogPositions, nil
}

func (r CatalogPositionsRepository) Get(ctx context.Context, filter bson.M) (models.CatalogPosition, error) {
	collection := r.Data.DB.Collection(catalogPositionsCollection)
	singleResult := collection.FindOne(ctx, filter)

	var catalogPosition models.CatalogPosition

	if err := singleResult.Decode(&catalogPosition); err != nil {
		return catalogPosition, err
	}

	return catalogPosition, nil
}

func (r CatalogPositionsRepository) Create(ctx context.Context, catalogPosition models.CatalogPosition) (models.CatalogPosition, error) {
	collection := r.Data.DB.Collection(catalogPositionsCollection)

	if _, err := collection.InsertOne(ctx, catalogPosition); err != nil {
		return catalogPosition, err
	}

	return catalogPosition, nil
}

func (r CatalogPositionsRepository) Update(ctx context.Context, filter bson.M, catalogPosition models.PartialCatalogPosition) (models.CatalogPosition, error) {
	collection := r.Data.DB.Collection(catalogPositionsCollection)
	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	singleResult := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": catalogPosition}, &updateOptions)

	var result models.CatalogPosition

	if err := singleResult.Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

func (r CatalogPositionsRepository) Delete(ctx context.Context, filter bson.M) error {
	collection := r.Data.DB.Collection(catalogPositionsCollection)

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (r CatalogPositionsRepository) CountDocuments(ctx context.Context, filter bson.M) (int64, error) {
	collection := r.Data.DB.Collection(catalogPositionsCollection)
	total, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return 0, err
	}

	return total, nil
}
