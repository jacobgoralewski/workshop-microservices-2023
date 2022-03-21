package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"

	"wrapper/todos"
)

type Handler struct {
	TodosClient *todos.Client
}

func (h *Handler) GetAllTodos(c echo.Context) error {
	err, allTodos := h.TodosClient.GetAllTodos()
	if err != nil {
		log.Err(err) // todo msg
		return err
	}

	return c.JSON(http.StatusOK, allTodos)
}
