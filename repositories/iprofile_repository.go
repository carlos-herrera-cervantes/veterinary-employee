package repositories

import (
	"context"
	"veterinary-employee/models"
)

//go:generate mockgen -destination=./mocks/iprofile_repository.go -package=mocks --build_flags=--mod=mod . IProfileRepository
type IProfileRepository interface {
	GetAll(ctx context.Context, page, pageSize int64) ([]models.Profile, error)
	Get(ctx context.Context, filter interface{}) (models.Profile, error)
	CountDocuments(ctx context.Context, filter interface{}) (int64, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Profile, error)
}
