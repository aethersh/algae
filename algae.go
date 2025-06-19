package main

import (
	"github.com/aethersh/algae/templates"
	"github.com/aethersh/algae/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app := fiber.New()

	app.Get("/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		c.Locals("name", name)
		if name == "" {
			name = "World"
		}
		return util.TemplRender(c, templates.Hello(name))
	})

	log.Fatal(app.Listen(":2152"))
}
