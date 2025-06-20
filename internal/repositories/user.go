package repositories

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
)

type UserRepository interface {
	Create(ctx context.Context, name, email string) (*domain.User, *errors.Error)
	GetOneById(ctx context.Context, id string) (*domain.User, *errors.Error)
	GetOneByEmail(ctx context.Context, email string) (*domain.User, *errors.Error)
	GetAll(ctx context.Context) ([]*domain.User, *errors.Error)
	Update(ctx context.Context, id string, name, email *string) *errors.Error
	Delete(ctx context.Context, id string) *errors.Error
}
