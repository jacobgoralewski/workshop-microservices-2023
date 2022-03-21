package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"contracts/users"

	"users/db"
	"users/handlers"
)

func main() {
	dbConn := db.Init()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	usersServer := handlers.NewServer(dbConn)
	grpcServer := grpc.NewServer()
	users.RegisterUsersServer(grpcServer, usersServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %s", err)
	}
}
