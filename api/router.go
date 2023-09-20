package api

import (
	"sync"

	"github.com/dfalgout/gofiber/dal"
	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Register(app *fiber.App, dal *dal.Queries) {
	once.Do(func() {
		apiRoutes := app.Group("/api")

		todosApi := NewTodosApi(dal)
		apiRoutes.Mount("/", todosApi.Register())
	})
}
