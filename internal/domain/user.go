package domain

type User struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Tasks []*Task `json:"tasks"`
}
