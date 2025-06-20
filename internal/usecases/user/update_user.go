package user

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/danielalmeidafarias/go-clean/internal/repositories"
)

type UpdateUserUseCase struct {
	userRepository repositories.UserRepository
}

func (uc *UpdateUserUseCase) Exec(ctx context.Context, id string, name, email *string) *errors.Error {
	return uc.userRepository.Update(ctx, id, name, email)
}
