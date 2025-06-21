package http

import (
	"github.com/danielalmeidafarias/go-clean/internal/usecases/user"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

type UserHandler struct {
	createUserUseCase *user.CreateUserUseCase
	getUserUseCase    *user.GetUserUseCase
	updateUserUseCase *user.UpdateUserUseCase
	deleteUserUseCase *user.DeleteUserUseCase
	validate          *validator.Validate
}

func NewUserHandler(
	createUserUseCase *user.CreateUserUseCase,
	getUserUseCase *user.GetUserUseCase,
	updateUserUseCase *user.UpdateUserUseCase,
	deleteUserUseCase *user.DeleteUserUseCase,
	validate *validator.Validate,
) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
		getUserUseCase:    getUserUseCase,
		updateUserUseCase: updateUserUseCase,
		deleteUserUseCase: deleteUserUseCase,
		validate:          validate,
	}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/", h.GetUser)
	user.Post("/", h.CreateUser)
	user.Put("/:id", h.UpdateUser)
	user.Delete("/:id", h.DeleteUser)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) {
	userID := c.Query("id")

	if err := h.validate.Var(userID, "required,uuid"); err != nil {
		c.Status(400).JSON(ValidationError(err))
		return
	}

	user, err := h.getUserUseCase.Exec(c.Context(), userID)
	if err != nil {
		c.Status(HttpStatusCode[err.Code]).JSON(err)
		return
	}

	c.Status(200).JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) {
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) {
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) {
}
