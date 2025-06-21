package http

import (
	"github.com/danielalmeidafarias/go-clean/internal/usecases/task"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
)

type TaskHandler struct {
	createTaskUseCase   *task.CreateTaskUseCase
	getTaskUseCase      *task.GetTaskUseCase
	updateTaskUseCase   *task.UpdateTaskUseCase
	deleteTaskUseCase   *task.DeleteTaskUseCase
	getUserTasksUseCase *task.GetUserTasks
	finishTaskUseCase   *task.FinishTaskUseCse
	unfinishTaskUseCase *task.UnfinishTaskUseCase
	changeOwnerUseCase  *task.ChangeOwnerUseCase
	validate            *validator.Validate
}

func NewTaskHandler(
	createTaskUseCase *task.CreateTaskUseCase,
	getTaskUseCase *task.GetTaskUseCase,
	updateTaskUseCase *task.UpdateTaskUseCase,
	deleteTaskUseCase *task.DeleteTaskUseCase,
	getUserTasksUseCase *task.GetUserTasks,
	finishTaskUseCase *task.FinishTaskUseCse,
	unfinishTaskUseCase *task.UnfinishTaskUseCase,
	changeOwnerUseCase *task.ChangeOwnerUseCase,
	validate *validator.Validate,
) *TaskHandler {
	return &TaskHandler{
		createTaskUseCase:   createTaskUseCase,
		getTaskUseCase:      getTaskUseCase,
		updateTaskUseCase:   updateTaskUseCase,
		deleteTaskUseCase:   deleteTaskUseCase,
		getUserTasksUseCase: getUserTasksUseCase,
		finishTaskUseCase:   finishTaskUseCase,
		unfinishTaskUseCase: unfinishTaskUseCase,
		changeOwnerUseCase:  changeOwnerUseCase,
		validate:            validate,
	}
}

func (h *TaskHandler) RegisterRoutes(app *fiber.App) {
	task := app.Group("/task")
	task.Get("/", h.GetTask)
	task.Post("/", h.CreateTask)
	task.Put("/:id", h.UpdateTask)
	task.Delete("/:id", h.DeleteTask)
	task.Get("/user/:userId", h.GetUserTasks)
	task.Post("/:id/finish", h.FinishTask)
	task.Post("/:id/unfinish", h.UnfinishTask)
	task.Post("/:id/change-owner", h.ChangeOwner)
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) {
	params := GetTaskDTO{
		Id: c.Query("id"),
	}

	if err := h.validate.Struct(params); err != nil {
		validationError(c, err)
		return
	}
	task, err := h.getTaskUseCase.Exec(c.Context(), params.Id)
	if err != nil {
		useCaseError(c, err)
		return
	}
	c.Status(200).JSON(task)
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) {
	var body CreateTaskDTO

	if err := c.BodyParser(&body); err != nil {
		invalidRequestBody(c)
		return
	}

	if err := h.validate.Struct(body); err != nil {
		validationError(c, err)
		return
	}
	task, err := h.createTaskUseCase.Exec(c.Context(), body.UserId, body.Name, body.Description)
	if err != nil {
		useCaseError(c, err)
		return
	}
	c.Status(201).JSON(task)
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) {
	var body UpdateTaskDTO

	if err := c.BodyParser(&body); err != nil {
		invalidRequestBody(c)
		return
	}

	id := c.Params("id")
	body.Id = id

	if err := h.validate.Struct(body); err != nil {
		validationError(c, err)
		return
	}

	if err := h.updateTaskUseCase.Exec(c.Context(), body.Id, body.Name, body.Description); err != nil {
		useCaseError(c, err)
		return
	}
	c.SendStatus(200)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) {
	params := DeleteTaskDTO{
		Id: c.Params("id"),
	}

	if err := h.validate.Struct(params); err != nil {
		validationError(c, err)
		return
	}

	if err := h.deleteTaskUseCase.Exec(c.Context(), params.Id); err != nil {
		useCaseError(c, err)
		return
	}
	c.SendStatus(200)
}

func (h *TaskHandler) GetUserTasks(c *fiber.Ctx) {
	params := GetUserTasksDTO{
		UserId: c.Params("userId"),
	}

	if err := h.validate.Struct(params); err != nil {
		validationError(c, err)
		return
	}
	tasks, err := h.getUserTasksUseCase.Exec(c.Context(), params.UserId)
	if err != nil {
		useCaseError(c, err)
		return
	}
	c.Status(200).JSON(tasks)
}

func (h *TaskHandler) FinishTask(c *fiber.Ctx) {
	id := c.Params("id")
	if err := h.finishTaskUseCase.Exec(c.Context(), id); err != nil {
		useCaseError(c, err)
		return
	}
	c.SendStatus(200)
}

func (h *TaskHandler) UnfinishTask(c *fiber.Ctx) {
	id := c.Params("id")
	if err := h.unfinishTaskUseCase.Exec(c.Context(), id); err != nil {
		useCaseError(c, err)
		return
	}
	c.SendStatus(200)
}

type ChangeOwnerDTO struct {
	UserId string `validate:"required,uuid"`
}

func (h *TaskHandler) ChangeOwner(c *fiber.Ctx) {
	id := c.Params("id")
	var body ChangeOwnerDTO
	if err := c.BodyParser(&body); err != nil {
		invalidRequestBody(c)
		return
	}
	if err := h.validate.Struct(body); err != nil {
		validationError(c, err)
		return
	}
	if err := h.changeOwnerUseCase.Exec(c.Context(), id, body.UserId); err != nil {
		useCaseError(c, err)
		return
	}
	c.SendStatus(200)
}
