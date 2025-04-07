package main

import (
	"fmt"
	"sensetion/go-fiber/config"
	"sensetion/go-fiber/internal/pages"
	"sensetion/go-fiber/pkg/database"
	"sensetion/go-fiber/pkg/logger"
	"sensetion/go-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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

	//Storage Init
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	pages.SetupRoutes(&pages.SetupRoutesDeps{
		App:    app,
		DBPool: dbpool,
		Logger: customLogger,
		Store:  store,
	})

	port := config.GetPort()
	customLogger.Info(fmt.Sprintf("Сервер запущен на порту %s", port))
	if err := app.Listen(port); err != nil {
		customLogger.Error("Ошибка при запуске сервера", "error", err)
	}
}
