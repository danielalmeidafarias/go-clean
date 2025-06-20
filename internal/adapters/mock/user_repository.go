package mock

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
)

type UserRepository struct {
	users []*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: []*domain.User{
			{
				Id:    "88cc4b04-02d0-4282-9b89-91fda89e56f0",
				Name:  "Usuário 1",
				Email: "email1@gmail.com",
				Tasks: nil,
			},
			{
				Id:    "b29630ce-d1a7-4627-8f37-c920a2a92872",
				Name:  "Usuário 2",
				Email: "email2@gmail.com",
				Tasks: nil,
			},
			{
				Id:    "7548c328-622b-407e-80a2-ba5870df780a",
				Name:  "Usuário 3",
				Email: "email3@gmail.com",
				Tasks: nil,
			},
		},
	}
}
func (r *UserRepository) GetOneById(ctx context.Context, id string) (*domain.User, *errors.Error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.NewError("not found", errors.NotFound)
}

func (r *UserRepository) GetOneByEmail(ctx context.Context, email string) (*domain.User, *errors.Error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.NewError("not found", errors.NotFound)
}

func (r *UserRepository) Create(ctx context.Context, name, email string) (*domain.User, *errors.Error) {
	if _, err := r.GetOneByEmail(ctx, email); err == nil {
		return nil, errors.NewError("unique constraint violation", errors.Conflict)
	}

	newUser := &domain.User{
		Id:    "8d38b17f-8460-4c2c-bed7-12ad6ec88538",
		Name:  name,
		Email: email,
		Tasks: nil,
	}

	r.users = append(r.users, newUser)

	return newUser, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, *errors.Error) {
	return r.users, nil
}

func (r *UserRepository) Update(ctx context.Context, id string, name, email *string) *errors.Error {
	user, err := r.GetOneById(ctx, id)
	if err != nil {
		return err
	}

	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) *errors.Error {
	if _, err := r.GetOneById(ctx, id); err != nil {
		return err
	}

	var usersCopy []*domain.User

	for _, user := range r.users {
		if user.Id != id {
			usersCopy = append(usersCopy, user)
		}
	}

	r.users = usersCopy
	return nil
}
