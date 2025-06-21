package http

type CreateTaskDTO struct {
	UserId      string `validate:"required,uuid"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

type GetTaskDTO struct {
	Id string `validate:"required,uuid"`
}

type UpdateTaskDTO struct {
	Id          string `validate:"required,uuid"`
	Name        *string
	Description *string
}

type DeleteTaskDTO struct {
	Id string `validate:"required,uuid"`
}

type GetUserTasksDTO struct {
	UserId string `validate:"required,uuid"`
}
