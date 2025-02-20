package home

import "github.com/gofiber/fiber/v2"

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
	h.router.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Welcome to the home page")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params is undefined")
}
