package mock

import (
	"context"

	"github.com/danielalmeidafarias/go-clean/internal/domain"
	"github.com/danielalmeidafarias/go-clean/internal/errors"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks []*domain.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: []*domain.Task{
			{
				Id:          "425f332b-49db-4e9b-990d-bc62b206890c",
				Name:        "Task 1",
				Description: "Description Task 1",
				Done:        false,
				UserId:      "88cc4b04-02d0-4282-9b89-91fda89e56f0",
			},
			{
				Id:          "481fb9f7-b9ea-428d-a7c5-8f580887c6f1",
				Name:        "Task 2",
				Description: "Description Task 2",
				Done:        false,
				UserId:      "b29630ce-d1a7-4627-8f37-c920a2a92872",
			},
			{
				Id:          "adea9980-01fc-4d68-a591-414dd50fa5a2",
				Name:        "Task 3",
				Description: "Description Task 3",
				Done:        false,
				UserId:      "7548c328-622b-407e-80a2-ba5870df780a",
			},
			{
				Id:          "57d093a0-5914-4d3a-91e5-c00ee2f70725",
				Name:        "Task 4",
				Description: "Description Task 4",
				Done:        false,
				UserId:      "88cc4b04-02d0-4282-9b89-91fda89e56f0",
			},
			{
				Id:          "f708a6e5-ae4b-452a-924f-63d2dac91a57",
				Name:        "Task 5",
				Description: "Description Task 5",
				Done:        false,
				UserId:      "b29630ce-d1a7-4627-8f37-c920a2a92872",
			},
		},
	}
}

func (r *TaskRepository) Create(ctx context.Context, name, description, userID string) (*domain.Task, *errors.Error) {
	newTask := &domain.Task{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		Done:        false,
		UserId:      userID,
	}

	r.tasks = append(r.tasks, newTask)

	return newTask, nil
}
func (r *TaskRepository) GetOneByID(ctx context.Context, id string) (*domain.Task, *errors.Error) {
	for _, task := range r.tasks {
		if task.Id == id {
			return task, nil
		}
	}

	return nil, errors.NewError("not found", errors.NotFound)
}
func (r *TaskRepository) GetAll(ctx context.Context) ([]*domain.Task, *errors.Error) {
	return r.tasks, nil
}
func (r *TaskRepository) GetByUser(ctx context.Context, userID string) ([]*domain.Task, *errors.Error) {
	var userTasks []*domain.Task

	for _, task := range r.tasks {
		if task.UserId == userID {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks, nil
}
func (r *TaskRepository) Update(ctx context.Context, id string, name, description, userID *string, done *bool) *errors.Error {
	task, err := r.GetOneByID(ctx, id)
	if err != nil {
		return err
	}

	if name != nil {
		task.Name = *name
	}

	if description != nil {
		task.Description = *description
	}

	if userID != nil {
		task.UserId = *userID
	}

	if done != nil {
		task.Done = *done
	}

	return nil
}
func (r *TaskRepository) Delete(ctx context.Context, id string) *errors.Error {
	if _, err := r.GetOneByID(ctx, id); err != nil {
		return err
	}

	var tasksCopy []*domain.Task

	for _, task := range r.tasks {
		if task.Id != id {
			tasksCopy = append(tasksCopy, task)
		}
	}

	r.tasks = tasksCopy
	return nil
}
