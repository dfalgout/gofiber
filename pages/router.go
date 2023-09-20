package pages

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var once sync.Once

func Register(app *fiber.App) {
	once.Do(func() {
		homeRoutes := NewHomeRoutes()
		todosRoutes := NewTodoRoutes()

		app.Mount("/", homeRoutes.Routes)
		app.Mount("/", todosRoutes.Routes)
	})
}
