package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type DeleteTaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewDeleteTaskUseCase(taskRepository repositories.TaskRepository) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (uc *DeleteTaskUseCase) Exec(ctx context.Context, taskID string) *errors.Error {
	errCtx := "error removing the task"

	_, err := uc.taskRepository.GetOneByID(ctx, taskID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	err = uc.taskRepository.Delete(ctx, taskID)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
