package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// ---- Init ----
	fmt.Println("Staring server...")
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone:   "UTC",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	// ---- Routes ----
	// Static files for GUI
	app.Static("/", "./www/public")

	// API
	api := app.Group("/api")
	api.Get("/helloworld", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// ---- Start ----
	app.Listen("127.0.0.1:3000")
}
