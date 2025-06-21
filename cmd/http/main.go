package main

import (
	"github.com/danielalmeidafarias/go-clean/internal/adapters/mock"
	"github.com/danielalmeidafarias/go-clean/internal/handlers/http"
	"github.com/danielalmeidafarias/go-clean/internal/usecases/user"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	userRepo := mock.NewUserRepository()

	createUserUC := user.NewCreateUserUseCase(userRepo)
	updateUserUC := user.NewUpdateUserUseCase(userRepo)
	getUserUC := user.NewGetUserUseCase(userRepo)
	deleteUserUC := user.NewDeleteUserUseCase(userRepo)

	httpServer := http.NewHttpServer(
		createUserUC,
		getUserUC,
		updateUserUC,
		deleteUserUC,
	)

	httpServer.Start()
}
