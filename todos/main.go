package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"todos/db"
	"todos/handlers"
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
	e.POST("/todos", h.CreateTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)

	e.Logger.Fatal(e.Start(":8081"))
}
