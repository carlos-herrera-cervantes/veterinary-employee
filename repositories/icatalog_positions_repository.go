package repositories

import (
	"context"

	"veterinary-employee/models"

	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockgen -destination=./mocks/icatalog_positions_repository.go -package=mocks --build_flags=--mod=mod . ICatalogPositionsRepository
type ICatalogPositionsRepository interface {
	GetAll(ctx context.Context, page, pageSize int64) ([]models.CatalogPosition, error)
	Get(ctx context.Context, filter bson.M) (models.CatalogPosition, error)
	Create(ctx context.Context, catalogPosition models.CatalogPosition) (models.CatalogPosition, error)
	Update(ctx context.Context, filter bson.M, catalogPosition models.PartialCatalogPosition) (models.CatalogPosition, error)
	Delete(ctx context.Context, filter bson.M) error
	CountDocuments(ctx context.Context, filter bson.M) (int64, error)
}
