package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Username string `json:"username"`
	Email    string `json:"email"`
}

type GetUserRequest struct {
	ID uint
}

type PostUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PutUserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
