package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type DeleteUserUseCase struct {
	userRepository repositories.UserRepository
}

func (uc *DeleteUserUseCase) Exec(ctx context.Context, id string) *errors.Error {
	errCtx := "erro removing the user"

	_, err := uc.userRepository.GetOneByID(ctx, id)
	if err != nil {
		return err.WithContext(errCtx)
	}

	return uc.userRepository.Delete(ctx, id)
}
