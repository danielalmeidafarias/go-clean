package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type UpdateTaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewUpdateTaskUseCase(taskRepository repositories.TaskRepository) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskRepository: taskRepository,
	}
}

func (uc *UpdateTaskUseCase) Exec(ctx context.Context, taskID string, name, description *string) *errors.Error {
	errCtx := "error updating the task"

	task, err := uc.taskRepository.GetOneByID(ctx, taskID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	if (name == nil && description == nil) || (name == &task.Name && description == &task.Description) {
		return errors.NewError("no changes required", errors.FailedPrecondition)
	}

	err = uc.taskRepository.Update(ctx, taskID, name, description, nil, nil)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
