package models

type Todo struct {
	ID          uint   `db:"id"          json:"id"`
	Description string `db:"description" json:"description"`
	UserID      uint   `db:"user_id"     json:"user_id"`
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
