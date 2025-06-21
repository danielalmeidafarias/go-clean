package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type UpdateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUpdateUserUseCase(userRepository repositories.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Exec(ctx context.Context, id string, name, email *string) *errors.Error {
	errCtx := "erros updating the user"

	user, err := uc.userRepository.GetOneByID(ctx, id)
	if err != nil {
		if err.Code == errors.NotFound {
			return errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}
		return errors.InternalError().WithContext(errCtx)
	}

	if name == nil && email == nil || name == &user.Name || email == &user.Email {
		return errors.NewError("no changes requested", errors.FailedPrecondition).WithContext(errCtx)
	}

	if email != nil {
		userWithEmail, err := uc.userRepository.GetOneByEmail(ctx, *email)
		if err != nil && err.Code != errors.NotFound {
			return errors.InternalError().WithContext(errCtx)
		}

		if userWithEmail != nil {
			return errors.NewError("email already in use", errors.Conflict).WithContext(errCtx)
		}
	}

	err = uc.userRepository.Update(ctx, id, name, email)
	if err != nil {
		return errors.InternalError().WithContext(errCtx)
	}

	return nil
}
