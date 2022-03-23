package todos

import (
	"github.com/pkg/errors"
	"wrapper/config"
	"wrapper/http_client"
	"wrapper/models"
)

type Client interface {
	GetAll() (error, []models.Todo)
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

func (sc HttpClient) GetAll() (error, []models.Todo) {
	var todos []models.Todo
	err := sc.httpClient.Get(sc.appConfig.TodosServiceURL, &todos)
	if err != nil {
		return errors.Wrap(err, "failed to fetch todos"), []models.Todo{}
	}

	return nil, todos
}
