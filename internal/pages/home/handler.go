package home

import (
	"log/slog"
	"sensetion/go-fiber/pkg/tadapter"
	"sensetion/go-fiber/views/pages/home"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
	logger *slog.Logger
}

func NewHandler(router fiber.Router, logger *slog.Logger) {
	h := &HomeHandler{
		router: router,
		logger: logger,
	}

	h.registerRoutes()
}

func (h *HomeHandler) registerRoutes() {
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := home.HomePage()

	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.logger.Error("Page not found", "url", c.OriginalURL())

	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}
