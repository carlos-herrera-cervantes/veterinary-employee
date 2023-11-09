package repositories

import (
	"context"
	"testing"

	"veterinary-employee/db"
	"veterinary-employee/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCatalogPositionsRepository_GetAll(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := catalogPositionsRepository.GetAll(ctxWithCancel, 0, 10)
		assert.Error(t, err)
	})
}

func TestCatalogPositionsRepository_Get(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := catalogPositionsRepository.Get(context.Background(), bson.M{"name": "test"})
		assert.Error(t, err)
	})
}

func TestCatalogPositionsRepository_Create(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := catalogPositionsRepository.Create(ctxWithCancel, models.CatalogPosition{})
		assert.Error(t, err)
	})
}

func TestCatalogPositionsRepository_Update(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := catalogPositionsRepository.Update(context.Background(), bson.M{"name": "test"}, models.PartialCatalogPosition{})
		assert.Error(t, err)
	})
}

func TestCatalogPositionsRepository_Delete(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return nil", func(t *testing.T) {
		err := catalogPositionsRepository.Delete(context.Background(), bson.M{"name": "test"})
		assert.Nil(t, err)
	})

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := catalogPositionsRepository.Delete(ctxWithCancel, bson.M{"name": "test"})
		assert.Error(t, err)
	})
}

func TestCatalogPositionsRepository_CountDocuments(t *testing.T) {
	catalogPositionsRepository := CatalogPositionsRepository{
		Data: db.New(),
	}

	t.Run("Should return 0 when counting documents fails", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		total, err := catalogPositionsRepository.CountDocuments(ctxWithCancel, bson.M{"name": "test"})
		assert.Error(t, err)
		assert.Zero(t, total)
	})

	t.Run("Should return 0", func(t *testing.T) {
		total, err := catalogPositionsRepository.CountDocuments(context.Background(), bson.M{"name": "test"})
		assert.Nil(t, err)
		assert.Zero(t, total)
	})
}
