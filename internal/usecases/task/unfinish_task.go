package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type UnfinishTaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewUnfinishTaskUseCase(
	taskRepository repositories.TaskRepository,
) *UnfinishTaskUseCase {
	return &UnfinishTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (uc *UnfinishTaskUseCase) Exec(ctx context.Context, taskID string) *errors.Error {
	errCtx := "error unfinishing the task owner"

	task, err := uc.taskRepository.GetOneByID(ctx, taskID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	if !task.Done {
		return errors.NewError("task is not finished", errors.FailedPrecondition).WithContext(errCtx)
	}

	done := false
	err = uc.taskRepository.Update(ctx, taskID, nil, nil, nil, &done)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
