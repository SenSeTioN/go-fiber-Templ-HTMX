package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sensetion/go-fiber/internal/pages"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	pages.SetupRoutes(app)

	app.Listen(":3000")
}
