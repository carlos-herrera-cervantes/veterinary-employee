package repositories

import (
	"context"
	"veterinary-employee/models"
)

type IAddressRepository interface {
	Get(ctx context.Context, filter interface{}) (models.Address, error)
	Create(ctx context.Context, address models.Address) (models.Address, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Address, error)
}
