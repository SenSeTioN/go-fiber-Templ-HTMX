package auth

import (
	"sensetion/go-fiber/pkg/tadapter"
	"sensetion/go-fiber/pkg/validator"
	"sensetion/go-fiber/views/entities/notification"
	"sensetion/go-fiber/views/pages/registration"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
func (h *AuthHandler) registrationPage(c *fiber.Ctx) error {
	component := registration.RegistrationPage()
	return tadapter.Render(c, component)
}

func (h *AuthHandler) registration(c *fiber.Ctx) error {
	form := AuthRegisterForm{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или некорректен",
		},
		&validators.StringIsPresent{
			Name:    "Name",
			Field:   form.Name,
			Message: "Имя не задано",
		},
		&validators.StringIsPresent{
			Name:    "Password",
			Field:   form.Password,
			Message: "Пароль не задан",
		},
	)

	var component templ.Component

	if len(errors.Errors) > 0 {
		component = notification.Notification(validator.FormatErrors(errors), notification.NotificationFail)
		return tadapter.Render(c, component)
	}

	component = notification.Notification("Регистрация успешна", notification.NotificationSuccess)
	return tadapter.Render(c, component)
}
