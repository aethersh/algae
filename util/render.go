package util

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func TemplRender(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}