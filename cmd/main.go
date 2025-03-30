package main

import (
	"fmt"
	"sensetion/go-fiber/config"
	"sensetion/go-fiber/internal/pages"
	"sensetion/go-fiber/pkg/database"
	"sensetion/go-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()
	dbConfig := config.NewDatabaseConfig()

	app := fiber.New()
	app.Use(recover.New())

	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app.Use(slogfiber.New(customLogger))
	app.Static("/public", "./public")

	dbpool := database.CreateDBPool(dbConfig, customLogger)
	defer dbpool.Close()

	pages.SetupRoutes(&pages.SetupRoutesArgs{
		App:    app,
		DBPool: dbpool,
		Logger: customLogger,
	})

	port := config.GetPort()
	customLogger.Info(fmt.Sprintf("Сервер запущен на порту %s", port))
	if err := app.Listen(port); err != nil {
		customLogger.Error("Ошибка при запуске сервера", "error", err)
	}
}
