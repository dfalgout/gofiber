package pages

import (
	"github.com/dfalgout/gofiber/render"
	"github.com/gofiber/fiber/v2"
)

type TodosRoutes struct {
	Routes *fiber.App
}

func NewTodoRoutes() *TodosRoutes {
	Routes := fiber.New()

	Routes.Get("/todos", getTodos)

	return &TodosRoutes{
		Routes,
	}
}

func getTodos(c *fiber.Ctx) error {
	return render.Templ(c, TodosPage())
}
