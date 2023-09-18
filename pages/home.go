package pages

import (
	"github.com/dfalgout/gofiber/render"
	"github.com/gofiber/fiber/v2"
)

type HomeRoutes struct {
	Routes *fiber.App
}

type CurrentPath struct {
	Path string
}

func NewHomeRoutes() *HomeRoutes {
	Routes := fiber.New()

	Routes.Get("/", home)

	return &HomeRoutes{
		Routes,
	}
}

func home(c *fiber.Ctx) error {
	return render.Templ(c, HomePage())
}
