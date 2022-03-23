package models

type Todo struct {
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        *User  `json:"user"`
}

type GetTodoRequest struct {
	ID uint `json:"id"`
}

type PostTodoRequest struct {
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type PutTodoRequest struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
