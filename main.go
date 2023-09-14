package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

//go:embed views/*
var views embed.FS

//go:embed assets/*
var assets embed.FS

type Todo struct {
	Name     string
	Complete bool
}

type NewTodoInput struct {
	Name     string
	Complete string
}

var todos = make([]*Todo, 0)

func main() {
	// serverConfig := config.NewServerConfig()
	engine := html.NewFileSystem(http.FS(views), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "assets")

	api := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{
			"Title": "Test",
		}, "views/layouts/main")
	})

	api.Get("/todos", func(c *fiber.Ctx) error {
		return c.Render("views/partials/todos/list", fiber.Map{
			"Todos": todos,
		})
	})

	api.Post("/todos", func(c *fiber.Ctx) error {
		newTodo := new(NewTodoInput)
		if err := c.BodyParser(newTodo); err != nil {
			return err
		}

		var complete = false
		if newTodo.Complete == "true" {
			complete = true
		}
		todo := &Todo{
			Name:     newTodo.Name,
			Complete: complete,
		}
		todos = append(todos, todo)

		return c.Render("views/partials/todos/single", fiber.Map{
			"Todo": todo,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
