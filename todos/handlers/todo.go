package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"todos/models"
)

type Handler struct {
	DB *gorm.DB // todo name?
}

func (h *Handler) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}
	log.Info().Msgf("got request for id: %d", id)

	var users models.Todo
	res := h.DB.First(&users, "id = ?", id) // todo repo
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	}
	if err = res.Error; err != nil {
		log.Err(err).Msg("failed to fetch from db")
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetTodos(c echo.Context) error {
	var users []models.Todo
	res := h.DB.Find(&users)
	if err := res.Error; err != nil {
		log.Err(err).Msg("failed to fetch from db")
		return err
	}
	if len(users) == 0 {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	request := models.Todo{}
	err := c.Bind(&request)
	if err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}
	log.Info().Msgf("got request: %+v", request)

	err = h.DB.Create(&request).Error
	if err != nil {
		log.Err(err).Msgf("failed to create user")
		return err
	}

	return c.JSON(http.StatusOK, request.ID)
}

func (h *Handler) UpdateTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}
	log.Info().Msgf("got request for id: %d", id)

	request := models.PutTodoRequest{}
	if err := c.Bind(&request); err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}

	err = h.DB.Model(models.Todo{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		log.Err(err).Msgf("failed to create user")
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("failed to parse request")
		return err
	}
	log.Info().Msgf("got request for id: %d", id)

	var users models.Todo
	res := h.DB.Delete(&users, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	}
	if err = res.Error; err != nil {
		log.Err(err).Msg("failed to fetch from db")
		return err
	}

	return c.NoContent(http.StatusOK)
}
