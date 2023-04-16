package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"wrapper/models"

	"github.com/golang/mock/gomock"

	"wrapper/mocks/todos"
	"wrapper/mocks/users"
)

func TestHandler_getAllTodos(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedTodosClient := todos.NewMockClient(ctrl)
	mockedUsersClient := users.NewMockClient(ctrl)

	mockedTodosClient.EXPECT().GetAll().Times(1).Return(
		[]models.Todo{
			{Description: "TEST_DESC", UserID: 111},
		},
		nil,
	)

	user := &models.User{
		Id:       1,
		Username: "TEST_USER_1",
		Email:    "TEST_USER_1@TEST.COM",
	}
	mockedUsersClient.EXPECT().GetUser(uint64(111)).Times(1).Return(
		user,
		nil,
	)

	h := NewHandler(mockedTodosClient, mockedUsersClient)

	allTodos, err := h.getAllTodos()

	assert.Nil(t, err)
	assert.Equal(t,
		[]models.Todo{
			{Description: "TEST_DESC", UserID: 111, User: user},
		},
		allTodos,
	)
}

func TestHandler_getAllTodos_2todos(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedTodosClient := todos.NewMockClient(ctrl)
	mockedUsersClient := users.NewMockClient(ctrl)

	mockedTodosClient.EXPECT().GetAll().Times(1).Return(
		[]models.Todo{
			{Description: "TEST_DESC_1", UserID: 111},
			{Description: "TEST_DESC_2", UserID: 222},
		},
		nil,
	)

	user := &models.User{
		Id:       1,
		Username: "TEST_USER_1",
		Email:    "TEST_USER_1@TEST.COM",
	}
	mockedUsersClient.EXPECT().GetUser(uint64(111)).Times(1).Return(
		user,
		nil,
	)
	mockedUsersClient.EXPECT().GetUser(uint64(222)).Times(1).Return(
		user,
		nil,
	)

	h := NewHandler(mockedTodosClient, mockedUsersClient)

	allTodos, err := h.getAllTodos()

	assert.Nil(t, err)
	assert.Equal(t,
		[]models.Todo{
			{Description: "TEST_DESC", UserID: 111, User: user},
		},
		allTodos,
	)
}
