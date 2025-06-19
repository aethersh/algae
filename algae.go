package main

import (
	"embed"
	"net/http"

	"github.com/aethersh/algae/templates"
	"github.com/aethersh/algae/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Embed a directory
//go:embed static/*
var embedDirStatic embed.FS

func main() {
	app := fiber.New()


	// CORS Configuration
	corsConfig, err := util.GenerateCORSConfig()
	if err != nil {
		log.Fatalf("Failed to generate CORS config: %v", err)
	}
	app.Use(cors.New(*corsConfig))

	// Static File Serving
	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(embedDirStatic),
		PathPrefix: "static",
		Browse: false,
	}))

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		sysinfo , _ := util.GetSystemInfo()
		return util.TemplRender(c, templates.HomePage(*sysinfo))
	})
	app.Get("/ping", func(c *fiber.Ctx) error {
		sysinfo , _ := util.GetSystemInfo()
		return util.TemplRender(c, templates.PingPage(*sysinfo))
	})
	app.Get("/traceroute", func(c *fiber.Ctx) error {
		sysinfo , _ := util.GetSystemInfo()
		return util.TemplRender(c, templates.TraceroutePage(*sysinfo))
	})
	app.Get("/bgp", func(c *fiber.Ctx) error {
		sysinfo , _ := util.GetSystemInfo()
		return util.TemplRender(c, templates.BGPPage(*sysinfo))
	})


	log.Fatal(app.Listen(":2152"))
}
