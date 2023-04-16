package main

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"users/db"
	"users/handlers"
	"users/proto"
)

const (
	port = 9000
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
	log.Info().Msgf("grpc server started at [::]:%d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %s", err)
	}
}
