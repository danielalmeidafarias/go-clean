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

func (uc *GetUserUseCase) Exec(ctx context.Context, id string) (*domain.User, *errors.Error) {
	return uc.userRepository.GetOneById(ctx, id)
}
