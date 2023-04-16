package users

import (
	"context"
	"wrapper/config"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"wrapper/proto"

	"wrapper/models"
)

//go:generate mockgen -destination=../mocks/users/users.go -package=users . Client
type Client interface {
	GetUser(uint64) (*models.User, error)
}

type GrpcClient struct {
	usersGrpcClient users.UsersClient
}

func NewServiceClient(appConfig config.AppConfig) *GrpcClient {
	conn, err := grpc.Dial(appConfig.UsersServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("did not connect: %s", err)
	}

	return &GrpcClient{
		usersGrpcClient: users.NewUsersClient(conn),
	}
}

func (sc GrpcClient) GetUser(id uint64) (*models.User, error) {
	user, err := sc.usersGrpcClient.GetUser(context.Background(), &users.GetUserRequest{
		Id: id,
	})
	if err != nil {
		log.Err(err).Msgf("failed to call: %s", err)
		return nil, err
	}

	return &models.User{
		Id:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
