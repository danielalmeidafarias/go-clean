package task

import (
	"context"
	"testing"

	mockRepo "github.com/danielalmeidafarias/go-clean/internal/adapters/mock"
)

func TestCreateTaskUseCase_Success(t *testing.T) {
	userRepo := mockRepo.NewUserRepository()
	taskRepo := mockRepo.NewTaskRepository()

	uc := NewCreateTaskUseCase(userRepo, taskRepo)
	ctx := context.Background()

	task, err := uc.Exec(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0", "Nova Task", "Descrição da Task")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task == nil || task.Name != "Nova Task" {
		t.Fatalf("expected task with name 'Nova Task', got %v", task)
	}
}

func TestCreateTaskUseCase_UserNotFound(t *testing.T) {
	userRepo := mockRepo.NewUserRepository()
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewCreateTaskUseCase(userRepo, taskRepo)
	ctx := context.Background()

	_, err := uc.Exec(ctx, "nao-existe", "Task", "Desc")
	if err == nil {
		t.Fatalf("expected error for user not found, got nil")
	}
}

func TestDeleteTaskUseCase_Success(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewDeleteTaskUseCase(taskRepo)
	ctx := context.Background()

	err := uc.Exec(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteTaskUseCase_NotFound(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewDeleteTaskUseCase(taskRepo)
	ctx := context.Background()

	err := uc.Exec(ctx, "nao-existe")
	if err == nil {
		t.Fatalf("expected error for task not found, got nil")
	}
}

func TestGetTaskUseCase_Success(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewGetTaskUseCase(taskRepo)
	ctx := context.Background()

	task, err := uc.Exec(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if err != nil || task == nil {
		t.Fatalf("expected task, got %v, %v", task, err)
	}
}

func TestGetTaskUseCase_NotFound(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewGetTaskUseCase(taskRepo)
	ctx := context.Background()

	_, err := uc.Exec(ctx, "nao-existe")
	if err == nil {
		t.Fatalf("expected error for task not found, got nil")
	}
}

func TestUpdateTaskUseCase_Success(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewUpdateTaskUseCase(taskRepo)
	ctx := context.Background()
	newName := "Task Atualizada"
	newDesc := "Desc Atualizada"

	err := uc.Exec(ctx, "425f332b-49db-4e9b-990d-bc62b206890c", &newName, &newDesc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, _ := taskRepo.GetOneByID(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if task.Name != newName || task.Description != newDesc {
		t.Fatalf("expected updated task, got %v", task)
	}
}

func TestUpdateTaskUseCase_NotFound(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewUpdateTaskUseCase(taskRepo)
	ctx := context.Background()
	newName := "Nome"

	err := uc.Exec(ctx, "nao-existe", &newName, nil)
	if err == nil {
		t.Fatalf("expected error for task not found, got nil")
	}
}

func TestGetUserTasks_Success(t *testing.T) {
	userRepo := mockRepo.NewUserRepository()
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewGetUserTasks(userRepo, taskRepo)
	ctx := context.Background()

	tasks, err := uc.Exec(ctx, "88cc4b04-02d0-4282-9b89-91fda89e56f0")
	if err != nil || len(tasks) == 0 {
		t.Fatalf("expected tasks, got %v, %v", tasks, err)
	}
}

func TestFinishTaskUseCase_Success(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewFinishTaskUseCse(taskRepo)
	ctx := context.Background()

	err := uc.Exec(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, _ := taskRepo.GetOneByID(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if !task.Done {
		t.Fatalf("expected task to be finished, got %v", task)
	}
}

func TestUnfinishTaskUseCase_Success(t *testing.T) {
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewUnfinishTaskUseCase(taskRepo)
	ctx := context.Background()

	err := uc.Exec(ctx, "f708a6e5-ae4b-452a-924f-63d2dac91a57")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, _ := taskRepo.GetOneByID(ctx, "f708a6e5-ae4b-452a-924f-63d2dac91a57")
	if task.Done {
		t.Fatalf("expected task to be unfinished, got %v", task)
	}
}

func TestChangeOwnerUseCase_Success(t *testing.T) {
	userRepo := mockRepo.NewUserRepository()
	taskRepo := mockRepo.NewTaskRepository()
	uc := NewChangeOwnerUseCase(userRepo, taskRepo)
	ctx := context.Background()
	newUserId := "b29630ce-d1a7-4627-8f37-c920a2a92872"

	err := uc.Exec(ctx, "425f332b-49db-4e9b-990d-bc62b206890c", newUserId)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, _ := taskRepo.GetOneByID(ctx, "425f332b-49db-4e9b-990d-bc62b206890c")
	if task.UserId != newUserId {
		t.Fatalf("expected owner changed, got %v", task)
	}
}
