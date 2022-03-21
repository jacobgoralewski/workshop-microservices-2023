package handlers

import (
	"context"
	"gorm.io/gorm"

	"contracts/users"
)

//go:generate protoc -I . --go_out=plugins=grpc:. some.proto // todo
type Server struct {
	DbConn *gorm.DB // todo name?
}

func (s Server) GetUser(ctx context.Context, request *users.GetUserRequest) (*users.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewServer(dbConn *gorm.DB) *Server {
	return &Server{DbConn: dbConn}
}
