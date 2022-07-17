package repositories

import (
	"context"
	"veterinary-employee/models"
)

//go:generate mockgen -destination=./mocks/iaddress_repository.go -package=mocks --build_flags=--mod=mod . IAddressRepository
type IAddressRepository interface {
	Get(ctx context.Context, filter interface{}) (models.Address, error)
	Create(ctx context.Context, address models.Address) (models.Address, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Address, error)
}
