package handlers

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"todos/models"
)

type Handler struct {
	DB *sqlx.DB
}

func (h *Handler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}
	log.Info().Msgf("got request for id: %d", id)

	var todo models.Todo
	err = h.DB.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		log.Err(err).Msgf("failed to fetch user with id=%d", id)
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *Handler) GetTodos(c echo.Context) error {
	var users []models.Todo
	err := h.DB.Select(&users, "SELECT * FROM todos")
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch from db")
		return err
	}
	if len(users) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}
