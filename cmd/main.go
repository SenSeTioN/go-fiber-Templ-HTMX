package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sensetion/go-fiber/config"
	"github.com/sensetion/go-fiber/internal/pages"
)

func main() {
	config.Init()

	app := fiber.New()
	app.Use(recover.New())

	pages.SetupRoutes(app)

	log.Println("Запуск сервера на :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
