package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type GetUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUserUseCase(userRepository repositories.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUserUseCase) Exec(ctx context.Context, id string) (*domain.User, *errors.Error) {
	errCtx := "error getting the user"

	user, err := uc.userRepository.GetOneByID(ctx, id)
	if err != nil {
		if err.Code == errors.NotFound {
			return nil, errors.NewError("user not found", errors.NotFound).WithContext(errCtx)
		}

		return nil, errors.InternalError().WithContext(errCtx)
	}

	return user, nil
}
