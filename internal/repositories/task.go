package repositories

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
)

type TaskRepository interface {
	Create(ctx context.Context, name, description, userID string) (*domain.Task, *errors.Error)
	GetOneByID(ctx context.Context, id string) (*domain.Task, *errors.Error)
	GetAll(ctx context.Context) ([]*domain.Task, *errors.Error)
	GetByUser(ctx context.Context, userID string) ([]*domain.Task, *errors.Error)
	Update(ctx context.Context, id string, name, description, userID *string, done *bool) *errors.Error
	Delete(ctx context.Context, id string) *errors.Error
}
