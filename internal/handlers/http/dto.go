package http

type CreateUserDTO struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type GetUserDTO struct {
	Id string `validate:"required,uuid"`
}

type UpdateUserDTO struct {
	Id    string `validate:"required,uuid"`
	Name  *string
	Email *string `validate:"omitempty,email"`
}

type DeleteUserDTO struct {
	Id string `validate:"required,uuid"`
}
