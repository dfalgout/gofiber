package render

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Templ(c *fiber.Ctx, component templ.Component) error {
	c.Response().Header.Add("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Context(), c)
}
