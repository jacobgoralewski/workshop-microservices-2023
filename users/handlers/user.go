package handlers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"contracts/users"

	"users/models"
)

type UsersServer struct {
	DbConn *gorm.DB // todo name?
}

func (us UsersServer) GetUser(_ context.Context, request *users.GetUserRequest) (*users.UserResponse, error) {
	var user models.User
	res := us.DbConn.First(&user, "id = ?", request.Id) // todo repo
	if err := res.Error; err != nil {
		log.Err(err).Msgf("failed to find user in database with id: %d", request.Id)
		return nil, errors.Wrapf(res.Error, "failed to find user in database with id: %d", request.Id)
	}
	x := &users.UserResponse{
		Id:       uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}
	return x, nil
}

func NewServer(dbConn *gorm.DB) *UsersServer {
	return &UsersServer{DbConn: dbConn}
}
