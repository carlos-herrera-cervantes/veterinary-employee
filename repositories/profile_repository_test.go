package repositories

import (
	"context"
	"testing"
	"time"

	"veterinary-employee/db"
	"veterinary-employee/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProfileRepositoryGetAll(t *testing.T) {
	mongoClient := db.New()
	profileRepository := ProfileRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := profileRepository.GetAll(ctxWithCancel, 0, 10)
		assert.Error(t, err)
	})

	t.Run("Should return an empty list", func(t *testing.T) {
		documents, err := profileRepository.GetAll(ctx, 0, 10)
		assert.NoError(t, err)
		assert.Empty(t, documents)
	})
}

func TestProfileRepositoryGet(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(profileCollection)
	profileRepository := ProfileRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := profileRepository.Get(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})

	t.Run("Should return a document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636ef5c55af4641dad2ef309")
		newProfile := models.Profile{
			EmployeeId:  objectId,
			Email:       "user@example.com",
			Name:        "Test",
			LastName:    "User",
			Gender:      "NotSpecified",
			PhoneNumber: "12345",
			Birthday:    time.Now(),
			Roles:       []string{"Employee"},
		}
		err := newProfile.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newProfile)
		assert.NoError(t, err)

		profile, err := profileRepository.Get(ctx, bson.M{"email": "user@example.com"})
		assert.NoError(t, err)
		assert.Equal(t, "Test", profile.Name)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestProfileRepositoryCountDocuments(t *testing.T) {
	mongoClient := db.New()
	profileRepository := ProfileRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return 0 when an error ocurrs", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		total, err := profileRepository.CountDocuments(ctxWithCancel, bson.M{})
		assert.Error(t, err)
		assert.Zero(t, total)
	})

	t.Run("Should return the total of documents", func(t *testing.T) {
		total, err := profileRepository.CountDocuments(ctx, bson.M{})
		assert.NoError(t, err)
		assert.Zero(t, total)
	})
}

func TestProfileRepositoryUpdate(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(profileCollection)
	profileRepository := ProfileRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := profileRepository.Update(ctxWithCancel, bson.M{}, models.PartialProfile{})
		assert.Error(t, err)
	})

	t.Run("Should update a document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636ef5c55af4641dad2ef309")
		newProfile := models.Profile{
			EmployeeId:  objectId,
			Email:       "user@example.com",
			Name:        "Test",
			LastName:    "User",
			Gender:      "NotSpecified",
			PhoneNumber: "12345",
			Birthday:    time.Now(),
			Roles:       []string{"Employee"},
		}
		err := newProfile.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newProfile)
		assert.NoError(t, err)

		partialProfile := models.PartialProfile{
			LastName: "Update",
		}
		updateResult, err := profileRepository.Update(ctx, bson.M{"email": "user@example.com"}, partialProfile)
		assert.NoError(t, err)
		assert.Equal(t, "Update", updateResult.LastName)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}
