package repositories

import (
	"context"
	"testing"

	"veterinary-employee/db"
	"veterinary-employee/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestAvatarRepositoryCountDocuments(t *testing.T) {
	mongoClient := db.New()
	avatarRepository := AvatarRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return 0 when an error occurs", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		total, err := avatarRepository.CountDocuments(ctxWithCancel, bson.M{})
		assert.Error(t, err)
		assert.Zero(t, total)
	})

	t.Run("Should return the total of documents", func(t *testing.T) {
		total, err := avatarRepository.CountDocuments(ctx, bson.M{})
		assert.NoError(t, err)
		assert.Zero(t, total)
	})
}

func TestAvatarRepositoryGet(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(avatarCollection)
	avatarRepository := AvatarRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error when document does not exist", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		_, err := avatarRepository.Get(ctx, bson.M{"employee_id": objectId})
		assert.Error(t, err)
	})

	t.Run("Should return document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAvatar := models.Avatar{
			EmployeeId: objectId,
			Path:       "profile.png",
		}
		err := newAvatar.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newAvatar)
		assert.NoError(t, err)

		singleResult, err := avatarRepository.Get(ctx, bson.M{"employee_id": objectId})
		assert.NoError(t, err)
		assert.Equal(t, "profile.png", singleResult.Path)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestAvatarRepositoryCreate(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(avatarCollection)
	avatarRepository := AvatarRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := avatarRepository.Create(ctxWithCancel, models.Avatar{})
		assert.Error(t, err)
	})

	t.Run("Should create a document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAvatar := models.Avatar{
			EmployeeId: objectId,
			Path:       "profile.png",
		}
		err := newAvatar.Validate()
		assert.NoError(t, err)

		_, err = avatarRepository.Create(ctx, newAvatar)
		assert.NoError(t, err)

		counter, err := collection.CountDocuments(ctx, bson.M{})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), counter)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestAvatarRepositoryUpdate(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(avatarCollection)
	avatarRepository := AvatarRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := avatarRepository.Update(ctxWithCancel, bson.M{}, models.Avatar{})
		assert.Error(t, err)
	})

	t.Run("Should update a document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAvatar := models.Avatar{
			EmployeeId: objectId,
			Path:       "profile.png",
		}
		err := newAvatar.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newAvatar)
		assert.NoError(t, err)

		partialAvatar := models.Avatar{
			Path: "new-profile.png",
		}
		_, err = avatarRepository.Update(ctx, bson.M{"employee_id": objectId}, partialAvatar)
		assert.NoError(t, err)

		singleResult := collection.FindOne(ctx, bson.M{"employee_id": objectId})
		avatar := models.Avatar{}

		err = singleResult.Decode(&avatar)
		assert.NoError(t, err)
		assert.Equal(t, "new-profile.png", avatar.Path)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestAvatarRepositoryDelete(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(avatarCollection)
	avatarRepository := AvatarRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		err := avatarRepository.Delete(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})

	t.Run("Should delete a document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAvatar := models.Avatar{
			EmployeeId: objectId,
			Path:       "profile.png",
		}
		err := newAvatar.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newAvatar)
		assert.NoError(t, err)

		err = avatarRepository.Delete(ctx, bson.M{"employee_id": objectId})
		assert.NoError(t, err)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}
