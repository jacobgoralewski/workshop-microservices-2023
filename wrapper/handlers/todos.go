package handlers

import (
	"github.com/pkg/errors"
	"net/http"
	"wrapper/models"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	"wrapper/todos"
	"wrapper/users"
)

type Handler struct {
	todosServiceClient todos.Client
	usersServiceClient users.Client
}

func NewHandler(todosServiceClient todos.Client, usersServiceClient users.Client) *Handler {
	return &Handler{
		todosServiceClient: todosServiceClient,
		usersServiceClient: usersServiceClient,
	}
}

func (h *Handler) GetAllTodos(c echo.Context) error {
	allTodos, err := h.getAllTodos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, allTodos)
}

func (h *Handler) getAllTodos() ([]models.Todo, error) {
	allTodos, err := h.todosServiceClient.GetAll()
	if err != nil {
		log.Err(err).Msg("failed to fetch all todos from todos service")
		return nil, errors.Wrap(err, "failed to fetch all todos from todos service")
	}

	for i, todo := range allTodos {
		user, err := h.usersServiceClient.GetUser(uint64(todo.UserID))
		if err != nil {
			log.Err(err).Msgf("failed to fetch user from users service with id: %d", todo.UserID)
		}

		allTodos[i].User = user
	}
	return allTodos, nil
}
