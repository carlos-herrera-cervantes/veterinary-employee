package repositories

import (
	"context"
	"veterinary-employee/models"
)

type IProfileRepository interface {
	GetAll(ctx context.Context, page, pageSize int64) ([]models.Profile, error)
	Get(ctx context.Context, filter interface{}) (models.Profile, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Profile, error)
}
