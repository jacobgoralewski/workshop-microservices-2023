package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wrapper/todos"

	"wrapper/handlers"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/", handlers.Home)

	h := &handlers.Handler{
		TodosClient: todos.NewClient(),
	}
	e.GET("/todos", h.GetAllTodos)

	e.Logger.Fatal(e.Start(":8082"))
}
