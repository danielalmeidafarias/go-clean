package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type CreateUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewCreateUserUseCase(userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Exec(ctx context.Context, name, email string) (*domain.User, *errors.Error) {
	errCtx := "error creating the user"

	user, err := uc.userRepository.GetOneByEmail(ctx, email)
	if err != nil && err.Code != errors.NotFound {
		return nil, err.WithContext(errCtx)
	}

	if user != nil {
		return nil, errors.NewError("email is already in use", errors.Conflict).WithContext(errCtx)
	}

	newUser, err := uc.userRepository.Create(ctx, name, email)
	if err != nil {
		return nil, err.WithContext(errCtx)
	}

	return newUser, nil
}
