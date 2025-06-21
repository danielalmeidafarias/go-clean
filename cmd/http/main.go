package main

import (
	"github.com/danielalmeidafarias/go-clean/internal/adapters/mock"
	"github.com/danielalmeidafarias/go-clean/internal/handlers/http"
	"github.com/danielalmeidafarias/go-clean/internal/usecases/task"
	"github.com/danielalmeidafarias/go-clean/internal/usecases/user"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	userRepo := mock.NewUserRepository()
	taskRepo := mock.NewTaskRepository()

	createUserUC := user.NewCreateUserUseCase(userRepo)
	updateUserUC := user.NewUpdateUserUseCase(userRepo)
	getUserUC := user.NewGetUserUseCase(userRepo)
	deleteUserUC := user.NewDeleteUserUseCase(userRepo)

	createTaskUC := task.NewCreateTaskUseCase(userRepo, taskRepo)
	getTaskUC := task.NewGetTaskUseCase(taskRepo)
	updateTaskUC := task.NewUpdateTaskUseCase(taskRepo)
	deleteTaskUC := task.NewDeleteTaskUseCase(taskRepo)
	getUserTasksUC := task.NewGetUserTasks(userRepo, taskRepo)
	finishTaskUC := task.NewFinishTaskUseCse(taskRepo)
	unfinishTaskUC := task.NewUnfinishTaskUseCase(taskRepo)
	changeOwnerUC := task.NewChangeOwnerUseCase(userRepo, taskRepo)

	httpServer := http.NewHttpServer(
		createUserUC,
		getUserUC,
		updateUserUC,
		deleteUserUC,
		createTaskUC,
		getTaskUC,
		updateTaskUC,
		deleteTaskUC,
		getUserTasksUC,
		finishTaskUC,
		unfinishTaskUC,
		changeOwnerUC,
	)

	httpServer.Start()
}
