package user

import (
	"context"
	"testing"

	mockRepo "github.com/danielalmeidafarias/go-clean/internal/adapters/mock"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
)

func TestCreateUserUseCase_Success(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := NewCreateUserUseCase(repo)
	ctx := context.Background()

	user, err := uc.Exec(ctx, "Novo Usuário", "novo@email.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if user == nil || user.Email != "novo@email.com" {
		t.Fatalf("expected user with email 'novo@email.com', got %v", user)
	}
}

func TestCreateUserUseCase_EmailConflict(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := NewCreateUserUseCase(repo)
	ctx := context.Background()

	_, err := uc.Exec(ctx, "Usuário 1", "email1@gmail.com")
	if err == nil || err.Code != errors.Conflict {
		t.Fatalf("expected conflict error, got %v", err)
	}
}

func TestDeleteUserUseCase_Success(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &DeleteUserUseCase{userRepository: repo}
	ctx := context.Background()

	err := uc.Exec(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteUserUseCase_NotFound(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &DeleteUserUseCase{userRepository: repo}
	ctx := context.Background()

	err := uc.Exec(ctx, "nao-existe")
	if err == nil || err.Code != errors.NotFound {
		t.Fatalf("expected not found error, got %v", err)
	}
}

func TestGetUserUseCase_Success(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &GetUserUseCase{userRepository: repo}
	ctx := context.Background()

	user, err := uc.Exec(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0")
	if err != nil || user == nil {
		t.Fatalf("expected user, got %v, %v", user, err)
	}
}

func TestGetUserUseCase_NotFound(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &GetUserUseCase{userRepository: repo}
	ctx := context.Background()

	user, err := uc.Exec(ctx, "nao-existe")
	if err == nil || err.Code != errors.NotFound || user != nil {
		t.Fatalf("expected not found error, got %v, %v", user, err)
	}
}

func TestUpdateUserUseCase_Success(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &UpdateUserUseCase{userRepository: repo}
	ctx := context.Background()
	newName := "Nome Atualizado"
	newEmail := "atualizado@email.com"

	err := uc.Exec(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0", &newName, &newEmail)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	user, _ := repo.GetOneByID(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0")
	if user.Name != newName || user.Email != newEmail {
		t.Fatalf("expected updated user, got %v", user)
	}
}

func TestUpdateUserUseCase_NotFound(t *testing.T) {
	repo := mockRepo.NewUserRepository()
	uc := &UpdateUserUseCase{userRepository: repo}
	ctx := context.Background()
	newName := "Nome"

	err := uc.Exec(ctx, "nao-existe", &newName, nil)
	if err == nil || err.Code != errors.NotFound {
		t.Fatalf("expected not found error, got %v", err)
	}
}
