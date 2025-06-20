package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type GetUserTasks struct {
	userRepository repositories.UserRepository
	taskRepository repositories.TaskRepository
}

func NewGetUserTasks(
	userRepository repositories.UserRepository,
	taskRepository repositories.TaskRepository,
) *GetUserTasks {
	return &GetUserTasks{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}

func (uc *GetUserTasks) Exec(ctx context.Context, userID string) ([]*domain.Task, *errors.Error) {
	errCtx := "error getting user's tasks"

	_, err := uc.userRepository.GetOneByID(ctx, userID)
	if err != nil {
		if err.Code == errors.NotFound {
			return nil, errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}

		return nil, errors.InternalError().WithContext(errCtx)
	}

	task, err := uc.taskRepository.GetByUser(ctx, userID)
	if err != nil {
		return nil, errors.InternalError().WithContext(errCtx)
	}

	return task, nil
}
