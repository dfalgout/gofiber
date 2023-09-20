package api

import (
	"log"

	"github.com/dfalgout/gofiber/ent"
	"github.com/dfalgout/gofiber/render"
	"github.com/gofiber/fiber/v2"
)

type TodosApi struct {
	routes *fiber.App
	client *ent.Client
}

func NewTodosApi(client *ent.Client) *TodosApi {
	routes := fiber.New()

	return &TodosApi{
		routes,
		client,
	}
}

func (ta *TodosApi) getTodos(c *fiber.Ctx) error {
	// user, ok := auth.GetUser(c)
	// if !ok {
	// 	return fmt.Errorf("unauthorized")
	// }
	// log.Println(user)
	todos, err := ta.client.Todo.Query().All(c.Context())
	if err != nil {
		return err
	}
	return render.Templ(c, TodosList(todos))
}

func (ta *TodosApi) createTodo(c *fiber.Ctx) error {
	newTodo := new(ent.Todo)
	if err := c.BodyParser(newTodo); err != nil {
		return err
	}
	todo, err := ta.client.Todo.Create().
		SetName(newTodo.Name).
		SetNillableCompleted(newTodo.Completed).
		Save(c.Context())
	if err != nil {
		log.Println(err)
		return err
	}
	return render.Templ(c, TodoSingle(todo))
}

func (ta *TodosApi) Register() *fiber.App {
	ta.routes.Get("/todos", ta.getTodos)
	ta.routes.Post("/todos", ta.createTodo)
	return ta.routes
}
