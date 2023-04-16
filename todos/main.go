package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"todos/db"
	"todos/handlers"
)

const (
	port = 8081
)

func main() {
	dbConn := db.Init()

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	h := handlers.Handler{
		DB: dbConn,
	}

	e.GET("/todos", h.GetTodos)
	e.GET("/todos/:id", h.GetTodo)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
