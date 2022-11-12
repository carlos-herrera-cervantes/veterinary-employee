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

func TestAddressRepositoryGet(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(addressCollection)
	addressRepository := AddressRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error when document does not exist", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		_, err := addressRepository.Get(ctx, bson.M{"employee_id": objectId})
		assert.Error(t, err)
	})

	t.Run("Should return document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		address := models.Address{
			EmployeeId:   objectId,
			Municipality: "Orizaba",
			PostalCode:   "39330",
			Street:       "Calle falsa 123",
			Colony:       "Bellas Artes",
			Number:       "25",
		}
		err := address.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, address)
		assert.NoError(t, err)

		singleResult, err := addressRepository.Get(ctx, bson.M{"employee_id": objectId})

		assert.NoError(t, err)
		assert.Equal(t, "Orizaba", singleResult.Municipality)
		assert.Equal(t, "39330", singleResult.PostalCode)
		assert.Equal(t, "Calle falsa 123", singleResult.Street)
		assert.Equal(t, "25", singleResult.Number)
		assert.Equal(t, "Bellas Artes", singleResult.Colony)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestAddressRepositoryCreate(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(addressCollection)
	addressRepository := AddressRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAddress := models.Address{
			EmployeeId:   objectId,
			Municipality: "Orizaba",
			PostalCode:   "39330",
			Street:       "Calle falsa 123",
			Colony:       "Bellas Artes",
			Number:       "25",
		}
		err := newAddress.Validate()
		assert.NoError(t, err)

		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err = addressRepository.Create(ctxWithCancel, newAddress)
		assert.Error(t, err)
	})

	t.Run("Should create document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAddress := models.Address{
			EmployeeId:   objectId,
			Municipality: "Orizaba",
			PostalCode:   "39330",
			Street:       "Calle falsa 123",
			Colony:       "Bellas Artes",
			Number:       "25",
		}
		err := newAddress.Validate()
		assert.NoError(t, err)

		_, err = addressRepository.Create(ctx, newAddress)
		assert.NoError(t, err)

		singleResult := collection.FindOne(ctx, bson.M{"employee_id": objectId})
		address := models.Address{}
		err = singleResult.Decode(&address)

		assert.NoError(t, err)
		assert.Equal(t, "Orizaba", address.Municipality)
		assert.Equal(t, "39330", address.PostalCode)
		assert.Equal(t, "Calle falsa 123", address.Street)
		assert.Equal(t, "25", address.Number)
		assert.Equal(t, "Bellas Artes", address.Colony)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}

func TestAddressRepositoryUpdate(t *testing.T) {
	mongoClient := db.New()
	collection := mongoClient.DB.Collection(addressCollection)
	addressRepository := AddressRepository{Data: mongoClient}
	ctx := context.Background()

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(ctx)
		cancel()

		_, err := addressRepository.Update(ctxWithCancel, bson.M{}, models.PartialAddress{})
		assert.Error(t, err)
	})

	t.Run("Should update document", func(t *testing.T) {
		objectId, _ := primitive.ObjectIDFromHex("636e7ba005f848dce0e4f80c")
		newAddress := models.Address{
			EmployeeId:   objectId,
			Municipality: "Orizaba",
			PostalCode:   "39330",
			Street:       "Calle falsa 123",
			Colony:       "Bellas Artes",
			Number:       "25",
		}
		err := newAddress.Validate()
		assert.NoError(t, err)

		_, err = collection.InsertOne(ctx, newAddress)
		assert.NoError(t, err)

		partialAddress := models.PartialAddress{
			Colony: "Cuicuilco",
		}

		_, err = addressRepository.Update(ctx, bson.M{"employee_id": objectId}, partialAddress)
		assert.NoError(t, err)

		singleResult := collection.FindOne(ctx, bson.M{"employee_id": objectId})
		address := models.Address{}

		err = singleResult.Decode(&address)
		assert.NoError(t, err)
		assert.Equal(t, "Cuicuilco", address.Colony)

		_, err = collection.DeleteMany(ctx, bson.M{})
		assert.NoError(t, err)
	})
}
