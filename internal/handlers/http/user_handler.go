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
	params := GetUserDTO{
		Id: c.Query("id"),
	}

	if err := h.validate.Struct(params); err != nil {
		validationError(c, err)
		return
	}

	user, err := h.getUserUseCase.Exec(c.Context(), params.Id)
	if err != nil {
		useCaseError(c, err)
		return
	}

	c.Status(200).JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) {
	var body CreateUserDTO

	if err := c.BodyParser(body); err != nil {
		invalidRequestBody(c)
		return
	}

	if err := h.validate.Struct(body); err != nil {
		validationError(c, err)
		return
	}

	user, err := h.createUserUseCase.Exec(c.Context(), body.Name, body.Email)
	if err != nil {
		useCaseError(c, err)
		return
	}

	c.Status(201).JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) {
	var body UpdateUserDTO

	if err := c.BodyParser(body); err != nil {
		invalidRequestBody(c)
		return
	}

	id := c.Params("id")
	body.Id = id

	if err := h.validate.Struct(body); err != nil {
		validationError(c, err)
		return
	}

	if err := h.updateUserUseCase.Exec(c.Context(), body.Id, body.Name, body.Email); err != nil {
		useCaseError(c, err)
		return
	}

	c.SendStatus(200)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) {
	params := DeleteUserDTO{
		Id: c.Params("id"),
	}

	if err := h.validate.Struct(params); err != nil {
		validationError(c, err)
		return
	}

	if err := h.deleteUserUseCase.Exec(c.Context(), params.Id); err != nil {
		useCaseError(c, err)
		return
	}

	c.SendStatus(200)
}
