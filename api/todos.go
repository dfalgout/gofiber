package api

import (
	"github.com/dfalgout/gofiber/render"
	"github.com/gofiber/fiber/v2"
)

type TodosApi struct {
	Routes *fiber.App
}

func NewTodosApi() *TodosApi {
	Routes := fiber.New()

	Routes.Get("/", getTodos)
	Routes.Post("/", createTodo)

	return &TodosApi{
		Routes,
	}
}

type Todo struct {
	Name     string
	Complete bool
}

type NewTodoInput struct {
	Name string
}

var todos = make([]*Todo, 0)

func getTodos(c *fiber.Ctx) error {
	return render.Templ(c, TodosList(todos))
}

func createTodo(c *fiber.Ctx) error {
	newTodo := new(NewTodoInput)
	if err := c.BodyParser(newTodo); err != nil {
		return err
	}
	todo := &Todo{
		Name: newTodo.Name,
	}
	todos = append(todos, todo)

	return render.Templ(c, TodoSingle(todo))
}
