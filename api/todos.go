package api

import (
	"log"

	"github.com/dfalgout/gofiber/dal"
	"github.com/dfalgout/gofiber/render"
	"github.com/gofiber/fiber/v2"
)

type TodosApi struct {
	routes *fiber.App
	dal    *dal.Queries
}

func NewTodosApi(dal *dal.Queries) *TodosApi {
	routes := fiber.New()

	return &TodosApi{
		routes,
		dal,
	}
}

type NewTodoInput struct {
	Name string
}

func (ta *TodosApi) getTodos(c *fiber.Ctx) error {
	todos, err := ta.dal.ListTodos(c.Context())
	if err != nil {
		return err
	}
	log.Println(todos)
	return render.Templ(c, TodosList(todos))
}

func (ta *TodosApi) createTodo(c *fiber.Ctx) error {
	newTodo := dal.CreateTodoParams{}
	if err := c.BodyParser(&newTodo); err != nil {
		return err
	}
	todo, err := ta.dal.CreateTodo(c.Context(), newTodo)
	if err != nil {
		return err
	}
	return render.Templ(c, TodoSingle(todo))
}

func (ta *TodosApi) Register() *fiber.App {
	ta.routes.Get("/todos", ta.getTodos)
	ta.routes.Post("/todos", ta.createTodo)
	return ta.routes
}
