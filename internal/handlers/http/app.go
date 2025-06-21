package http

import (
	"fmt"
	"log"
	"os"

	"github.com/danielalmeidafarias/go-clean/internal/usecases/user"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

type HttpServer struct {
	app         *fiber.App
	userHandler *UserHandler
}

func NewHttpServer(
	createUserUC *user.CreateUserUseCase,
	getUserUC *user.GetUserUseCase,
	updateUserUC *user.UpdateUserUseCase,
	deleteUserUC *user.DeleteUserUseCase,
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

	return &HttpServer{
		app:         app,
		userHandler: userHandler,
	}
}

func (s *HttpServer) Start() {
	s.userHandler.RegisterRoutes(s.app)

	httpPort := os.Getenv("HTTP_PORT")

	log.Fatal(s.app.Listen(fmt.Sprintf(":%s", httpPort)))
}
