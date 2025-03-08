package main

import (
	"sensetion/go-fiber/config"
	"sensetion/go-fiber/internal/pages"
	"sensetion/go-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()

	engine := html.New("./html", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(recover.New())

	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	app.Use(slogfiber.New(customLogger))

	pages.SetupRoutes(app)

	customLogger.Info("Запуск сервера на :4242")
	if err := app.Listen(":4242"); err != nil {
		customLogger.Error("Ошибка при запуске сервера", "error", err)
	}
}
