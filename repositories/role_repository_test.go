package repositories

import (
	"context"
	"testing"

	"veterinary-employee/db"
	"veterinary-employee/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRoleRepositoryGetAll(t *testing.T) {
	mongoClient := db.New()
	roleRepository := RoleRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := roleRepository.GetAll(ctxWithCancel)
		assert.Error(t, err)
	})

	t.Run("Should return an empty list", func(t *testing.T) {
		documents, err := roleRepository.GetAll(ctx)
		assert.NoError(t, err)
		assert.Empty(t, documents)
	})
}

func TestRoleRepositoryGet(t *testing.T) {
	mongoClient := db.New()
	roleRepository := RoleRepository{Data: mongoClient}
	ctx := context.Background()
	collection := mongoClient.DB.Collection(roleCollection)

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := roleRepository.Get(ctxWithCancel, bson.M{})
		assert.Error(t, err)
	})

	t.Run("Should return a document", func(t *testing.T) {
		newRole := models.Role{
			Name:   "Employee",
			Active: true,
		}
		err := newRole.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newRole)
		assert.NoError(t, err)

		_, err = roleRepository.Get(ctx, bson.M{"name": "Employee"})
		assert.NoError(t, err)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestRoleRepositoryCreate(t *testing.T) {
	mongoClient := db.New()
	roleRepository := RoleRepository{Data: mongoClient}
	ctx := context.Background()
	collection := mongoClient.DB.Collection(roleCollection)

	t.Run("Should return an error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := roleRepository.Create(ctxWithCancel, models.Role{})
		assert.Error(t, err)
	})

	t.Run("Should create a document", func(t *testing.T) {
		newRole := models.Role{
			Name:   "Employee",
			Active: true,
		}
		err := newRole.Validate()
		assert.NoError(t, err)

		_, err = roleRepository.Create(ctx, newRole)
		assert.NoError(t, err)

		counter, err := collection.CountDocuments(ctx, bson.M{})
		assert.NoError(t, err)
		assert.Equal(t, int64(1), counter)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestRoleRepositoryUpdate(t *testing.T) {
	mongoClient := db.New()
	roleRepository := RoleRepository{Data: mongoClient}
	ctx := context.Background()
	collection := mongoClient.DB.Collection(roleCollection)

	t.Run("Should return an error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := roleRepository.Update(ctxWithCancel, bson.M{}, models.Role{})
		assert.Error(t, err)
	})

	t.Run("Should update a document", func(t *testing.T) {
		newRole := models.Role{
			Name:   "Employee",
			Active: true,
		}
		err := newRole.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newRole)
		assert.NoError(t, err)

		active := false
		partialRole := models.PartialRole{Active: &active}
		updateResult, err := roleRepository.Update(ctx, bson.M{"name": "Employee"}, partialRole)
		assert.NoError(t, err)
		assert.False(t, updateResult.Active)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}
