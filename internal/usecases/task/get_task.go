package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type GetTaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewGetTaskUseCase(taskRepository repositories.TaskRepository) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (uc *GetTaskUseCase) Exec(ctx context.Context, id string) (*domain.Task, *errors.Error) {
	errCtx := "error finding the task"

	task, err := uc.taskRepository.GetOneByID(ctx, id)
	if err != nil {
		if err.Code == errors.NotFound {
			return nil, errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return nil, errors.InternalError().WithContext(errCtx)
	}

	return task, nil
}
