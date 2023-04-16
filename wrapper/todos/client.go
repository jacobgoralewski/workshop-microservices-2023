package todos

import (
	"github.com/pkg/errors"
	"wrapper/config"
	"wrapper/http_client"
	"wrapper/models"
)

//go:generate mockgen -destination=../mocks/todos/todos.go -package=todos . Client
type Client interface {
	GetAll() ([]models.Todo, error)
}

type HttpClient struct {
	appConfig  config.AppConfig
	httpClient *http_client.HttpClient
}

func NewServiceClient(appConfig config.AppConfig, httpClient *http_client.HttpClient) *HttpClient {
	return &HttpClient{
		httpClient: httpClient,
		appConfig:  appConfig,
	}
}

func (sc HttpClient) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	err := sc.httpClient.Get(sc.appConfig.TodosServiceURL, &todos)
	if err != nil {
		return []models.Todo{}, errors.Wrap(err, "failed to fetch todos")
	}

	return todos, nil
}
