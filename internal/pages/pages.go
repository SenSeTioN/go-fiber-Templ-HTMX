package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sensetion/go-fiber/internal/pages/home"
)

func SetupRoutes(app *fiber.App) {
	v := app.Group("/v1")
	api := v.Group("/api")

	home.NewHandler(api)
}
