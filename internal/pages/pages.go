package pages

import (
	"sensetion/go-fiber/internal/pages/auth"
	"sensetion/go-fiber/internal/pages/home"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	home.NewHandler(app)
	auth.NewAuthHandler(app)
}
