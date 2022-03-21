package todos

//import (
//	"context"
//
//	"github.com/rs/zerolog/log"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//
//	"contracts/todos"
//)
//
//type Client struct { // todo interface, name?
//	todosGrpcClient todos.TodosClient
//}
//
//func NewClient() *Client {
//	// todo config
//	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatal().Msgf("did not connect: %s", err)
//	}
//
//	c := todos.NewTodosClient(conn)
//
//	return &Client{
//		todosGrpcClient: c,
//	}
//}
//
//func (c Client) GetAllTodos() (error, []*todos.TodoResponse) {
//	allTodos, err := c.todosGrpcClient.GetAllTodos(context.Background(), &todos.GetAllTodosRequest{})
//	if err != nil {
//		log.Err(err).Msgf("Error when calling SayHello: %s", err) // todo msg
//		return err, nil
//	}
//	log.Info().Msgf("%+v", allTodos)
//
//	return nil, allTodos.Todos
//}
