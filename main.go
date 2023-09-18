package main

import (
	"embed"
	"log"

	"github.com/dfalgout/gofiber/api"
	"github.com/dfalgout/gofiber/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed assets/*
var assets embed.FS

type Route struct {
	Name string
	Path string
}

var loggedIn = false

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())

	app.Static("/assets", "assets")

	// view routes
	homeRoutes := pages.NewHomeRoutes()
	todosRoutes := pages.NewTodoRoutes()

	app.Mount("/", homeRoutes.Routes)
	app.Mount("/", todosRoutes.Routes)

	// api routes
	apiRoutes := app.Group("/api")

	todosApi := api.NewTodosApi()
	apiRoutes.Mount("/todos", todosApi.Routes)

	log.Fatal(app.Listen(":3000"))
}
