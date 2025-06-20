package task

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type ChangeOwnerUseCase struct {
	userRepository repositories.UserRepository
	taskRepository repositories.TaskRepository
}

func NewChangeOwnerUseCase(
	userRepository repositories.UserRepository,
	taskRepository repositories.TaskRepository,
) *ChangeOwnerUseCase {
	return &ChangeOwnerUseCase{
		userRepository: userRepository,
		taskRepository: taskRepository,
	}
}

func (uc *ChangeOwnerUseCase) Exec(ctx context.Context, taskID string, userID string) *errors.Error {
	errCtx := "error changing the task owner"

	_, err := uc.taskRepository.GetOneByID(ctx, taskID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("task not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	_, err = uc.userRepository.GetOneByID(ctx, userID)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}

		return errors.InternalError().WithContext(errCtx)
	}

	err = uc.taskRepository.Update(ctx, nil, nil, &userID, nil)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
