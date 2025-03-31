package home

import (
	"sensetion/go-fiber/pkg/tadapter"
	"sensetion/go-fiber/views/pages/home"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}

	h.registerRoutes()
}

func (h *HomeHandler) registerRoutes() {
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := home.HomePage()

	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {

	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}
