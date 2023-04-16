package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wrapper/config"
	"wrapper/http_client"
	"wrapper/todos"
	"wrapper/users"

	"wrapper/handlers"
)

func main() {
	appConfig := config.GetConfig()

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	h := handlers.NewHandler(
		todos.NewServiceClient(
			appConfig,
			http_client.NewHttpClient(),
		),
		users.NewServiceClient(
			appConfig,
		),
	)

	e.GET("/todos", h.GetAllTodos)

	e.Logger.Fatal(e.Start(":8082"))
}
