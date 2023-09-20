package api

import (
	"sync"

	"github.com/dfalgout/gofiber/ent"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Register(app *fiber.App, client *ent.Client) {
	once.Do(func() {
		apiRoutes := app.Group("/api")

		todosApi := NewTodosApi(client)
		apiRoutes.Mount("/", todosApi.Register())
	})
}
