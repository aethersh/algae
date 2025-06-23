package main

import (
	"embed"
	"net/http"

	"github.com/aethersh/algae/mtr"
	"github.com/aethersh/algae/templates"
	"github.com/aethersh/algae/util"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static/*
var embedDirStatic embed.FS

func main() {
	app := fiber.New()

	// Request Logging
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &util.RequestLogger,
	}))

	// CORS Configuration
	corsConfig, err := util.GenerateCORSConfig()
	if err != nil {
		log.Fatalf("Failed to generate CORS config: %v", err)
	}
	app.Use(cors.New(*corsConfig))

	// Static File Serving
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse:     false,
	}))

	// Init helpers/data
	sysinfo, _ := util.GetSystemInfo()

	// ROUTES
	// Views
	app.Get("/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		c.Locals("name", name)

		component := templates.HomePage(*sysinfo)
		if name == "ping" {
			component = templates.PingPage(*sysinfo)
		} else if name == "traceroute" {
			component = templates.TraceroutePage(*sysinfo)
		} else if name == "bgp" {
			component = templates.BGPPage(*sysinfo)
		}

		return util.TemplRender(c, component)
	})

	// Action Handlers
	app.Post("/ping", func(c *fiber.Ctx) error {
		host := c.FormValue("ipAddr")
		out, err := mtr.RunPingCmd(host)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.SendString(err.Error())
		}

		println(*out)

		component := templates.PingOutput(*out)
		return util.TemplRender(c, component)
	})

	log.Fatal(app.Listen(":2152"))
}
