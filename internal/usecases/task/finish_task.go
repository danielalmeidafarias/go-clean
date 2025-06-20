package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type FinishTaskUseCse struct {
	taskRepository repositories.TaskRepository
}

func NewFinishTaskUseCse(
	taskRepository repositories.TaskRepository,
) *FinishTaskUseCse {
	return &FinishTaskUseCse{
		taskRepository: taskRepository,
	}
}

func (uc *FinishTaskUseCse) Exec(ctx context.Context, taskID string, userID string) *errors.Error {
	errCtx := "error finishing the task owner"

	task, err := uc.taskRepository.GetOneByID(ctx, taskID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	if task.Done {
		return errors.NewError("task already finished", errors.FailedPrecondition).WithContext(errCtx)
	}

	done := true
	err = uc.taskRepository.Update(ctx, taskID, nil, nil, nil, &done)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
