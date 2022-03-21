package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	gorm.Model
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
