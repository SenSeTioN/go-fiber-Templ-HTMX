package auth

import (
	"sensetion/go-fiber/pkg/tadapter"
	"sensetion/go-fiber/views/pages/registration"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	router fiber.Router
}

func NewAuthHandler(router fiber.Router) {
	h := &AuthHandler{
		router: router,
	}

	h.registerRoutes()
}

func (h *AuthHandler) registerRoutes() {
	apiRouter := h.router.Group("/v1/api/auth")
	authRouter := h.router.Group("/auth")

	apiRouter.Post("/register", h.registration)
	authRouter.Get("/registration", h.registrationPage)
}

func (h *AuthHandler) registration(c *fiber.Ctx) error {
	return c.SendString("Registration")
}

func (h *AuthHandler) registrationPage(c *fiber.Ctx) error {
	component := registration.RegistrationPage()
	return tadapter.Render(c, component)
}
