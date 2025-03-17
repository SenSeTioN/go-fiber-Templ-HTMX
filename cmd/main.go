package main

import (
	"sensetion/go-fiber/config"
	"sensetion/go-fiber/internal/pages"
	"sensetion/go-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	app := fiber.New()
	app.Use(recover.New())

	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app.Use(slogfiber.New(customLogger))
	app.Static("/public", "./public")

	pages.SetupRoutes(app)

	customLogger.Info("Запуск сервера на :4242")
	if err := app.Listen(":4242"); err != nil {
		customLogger.Error("Ошибка при запуске сервера", "error", err)
	}
}
