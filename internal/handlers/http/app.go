package http

import (
	"fmt"
	"log"
	"os"

	"github.com/danielalmeidafarias/go-clean/internal/usecases/task"
	"github.com/danielalmeidafarias/go-clean/internal/usecases/user"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

type HttpServer struct {
	app          *fiber.App
	userHandler  *UserHandler
	tasksHandler *TaskHandler
}

func NewHttpServer(
	createUserUC *user.CreateUserUseCase,
	getUserUC *user.GetUserUseCase,
	updateUserUC *user.UpdateUserUseCase,
	deleteUserUC *user.DeleteUserUseCase,
	createTaskUC *task.CreateTaskUseCase,
	getTaskUC *task.GetTaskUseCase,
	updateTaskUC *task.UpdateTaskUseCase,
	deleteTaskUC *task.DeleteTaskUseCase,
	getUsersTasksTC *task.GetUserTasks,
	finishTaskUC *task.FinishTaskUseCse,
	unfinishTaskUC *task.UnfinishTaskUseCase,
	changeOwnerUC *task.ChangeOwnerUseCase,
) *HttpServer {
	app := fiber.New()
	validate := validator.New()

	userHandler := NewUserHandler(
		createUserUC,
		getUserUC,
		updateUserUC,
		deleteUserUC,
		validate,
	)

	taskHandler := NewTaskHandler(
		createTaskUC,
		getTaskUC,
		updateTaskUC,
		deleteTaskUC,
		getUsersTasksTC,
		finishTaskUC,
		unfinishTaskUC,
		changeOwnerUC,
		validate,
	)

	return &HttpServer{
		app:          app,
		userHandler:  userHandler,
		tasksHandler: taskHandler,
	}
}

func (s *HttpServer) Start() {
	s.userHandler.RegisterRoutes(s.app)
	s.tasksHandler.RegisterRoutes(s.app)

	httpPort := os.Getenv("HTTP_PORT")

	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", httpPort)))
}
