package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type CreateTaskUseCase struct {
	userRepository repositories.UserRepository
	taskRepository repositories.TaskRepository
}

func NewCreateTaskUseCase(
	userRepository repositories.UserRepository,
	taskRepository repositories.TaskRepository,
) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}

func (uc *CreateTaskUseCase) Exec(ctx context.Context, userID, name, description string) (*domain.Task, *errors.Error) {
	errCtx := "error creating the task"

	_, err := uc.userRepository.GetOneByID(ctx, userID)
	if err != nil {
		if err.Code == errors.NotFound {
			return nil, errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}

		return nil, errors.InternalError().WithContext(errCtx)
	}

	task, err := uc.taskRepository.Create(ctx, name, description, userID)
	if err != nil {
		return nil, errors.InternalError().WithContext(errCtx)
	}

	return task, nil
}
