package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"trivy-service/internal"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	// app.Use(compression.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
	app.Get("/scan", internal.GetVulnerability)
	// controllers.GetVulnerability()
	app.Listen(":4000")
}
