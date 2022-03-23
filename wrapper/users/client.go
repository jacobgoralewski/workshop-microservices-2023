package users

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"contracts/users"

	"wrapper/models"
)

type Client interface {
	GetUser(uint64) (error, *models.User)
}

type GrpcClient struct { // todo interface, name?
	usersGrpcClient users.UsersClient
}

func NewServiceClient() *GrpcClient {
	// todo config
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("did not connect: %s", err) // todo
	}

	return &GrpcClient{
		usersGrpcClient: users.NewUsersClient(conn),
	}
}

func (sc GrpcClient) GetUser(id uint64) (error, *models.User) {
	user, err := sc.usersGrpcClient.GetUser(context.Background(), &users.GetUserRequest{
		Id: id,
	})
	if err != nil {
		log.Err(err).Msgf("Error when calling SayHello: %s", err) // todo msg
		return err, nil
	}

	return nil, &models.User{
		Id:       uint(user.Id),
		Username: user.Username,
		Email:    user.Email,
	}
}
