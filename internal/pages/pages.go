package pages

import (
	"log/slog"
	"sensetion/go-fiber/internal/pages/home"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, logger *slog.Logger) {
	// v := app.Group("/v1")
	// api := v.Group("/api")

	home.NewHandler(app, logger)
}
