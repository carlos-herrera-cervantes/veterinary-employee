package repositories

import (
	"context"
	"veterinary-employee/models"
)

//go:generate mockgen -destination=./mocks/irole_repository.go -package=mocks --build_flags=--mod=mod . IRoleRepository
type IRoleRepository interface {
	GetAll(ctx context.Context) ([]models.Role, error)
	Get(ctx context.Context, filter interface{}) (models.Role, error)
	Create(ctx context.Context, role models.Role) (models.Role, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Role, error)
}
