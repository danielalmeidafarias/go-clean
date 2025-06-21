package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type DeleteUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewDeleteUserUseCase(userRepository repositories.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Exec(ctx context.Context, id string) *errors.Error {
	errCtx := "erro removing the user"

	_, err := uc.userRepository.GetOneByID(ctx, id)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}
		return errors.InternalError().WithContext(errCtx)
	}

	err = uc.userRepository.Delete(ctx, id)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
