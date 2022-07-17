package repositories

import (
	"context"
	"veterinary-employee/models"
)

//go:generate mockgen -destination=./mocks/iavatar_repository.go -package=mocks --build_flags=--mod=mod . IAvatarRepository
type IAvatarRepository interface {
	Get(ctx context.Context, filter interface{}) (models.Avatar, error)
	Create(ctx context.Context, avatar models.Avatar) (models.Avatar, error)
	Update(ctx context.Context, filter interface{}, document interface{}) (models.Avatar, error)
	Delete(ctx context.Context, filter interface{}) error
}
