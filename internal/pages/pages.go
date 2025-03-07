package pages

import (
	"sensetion/go-fiber/internal/pages/home"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v := app.Group("/v1")
	api := v.Group("/api")

	home.NewHandler(api)
}
