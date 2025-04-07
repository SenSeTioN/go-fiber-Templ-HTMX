package auth

import (
	"fmt"
	"sensetion/go-fiber/internal/types"
	"sensetion/go-fiber/internal/users"
	"sensetion/go-fiber/pkg/tadapter"
	"sensetion/go-fiber/pkg/validator"
	"sensetion/go-fiber/views/entities/notification"
	"sensetion/go-fiber/views/pages/login"
	"sensetion/go-fiber/views/pages/registration"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthHandlerDeps struct {
	router     fiber.Router
	repository *users.UserRepository
	store      *session.Store
}

func NewAuthHandler(router fiber.Router, repository *users.UserRepository, store *session.Store) {
	h := &AuthHandlerDeps{
		router:     router,
		repository: repository,
		store:      store,
	}

	h.registerRoutes()
}

func (h *AuthHandlerDeps) registerRoutes() {
	apiRouter := h.router.Group("/v1/api/auth")
	apiRouter.Post("/register", h.registration)
	apiRouter.Post("/login", h.login)
	apiRouter.Get("/logout", h.logout)

	authRouter := h.router.Group("/auth")
	authRouter.Get("/registration", h.registrationPage)
	authRouter.Get("/login", h.loginPage)
}

func (h *AuthHandlerDeps) registrationPage(c *fiber.Ctx) error {
	component := registration.RegistrationPage()
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *AuthHandlerDeps) registration(c *fiber.Ctx) error {
	form := types.AuthRegisterForm{
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
		return tadapter.Render(c, component, fiber.StatusBadRequest)
	}

	err := h.repository.CreateUser(form)
	if err != nil {
		h.repository.CustomLogger.Error("ошибка создания пользователя", "error", err)
		component = notification.Notification(err.Error(), notification.NotificationFail)
		return tadapter.Render(c, component, fiber.StatusBadRequest)
	}

	component = notification.Notification(fmt.Sprintf("Пользователь %s успешно зарегистрирован", form.Email), notification.NotificationSuccess)
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *AuthHandlerDeps) loginPage(c *fiber.Ctx) error {
	component := login.LoginPage()
	return tadapter.Render(c, component, fiber.StatusOK)
}

func (h *AuthHandlerDeps) login(c *fiber.Ctx) error {
	form := types.AuthLoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	user, err := h.repository.GetUserByCredentials(form.Email, form.Password)
	if err != nil {
		h.repository.CustomLogger.Error("ошибка получения пользователя", "error", err)
		component := notification.Notification(err.Error(), notification.NotificationFail)
		return tadapter.Render(c, component, fiber.StatusBadRequest)
	}

	if user != nil {
		session, err := h.store.Get(c)
		if err != nil {
			panic(err)
		}

		session.Set("email", user.Email)
		session.Set("name", user.Name)
		if err := session.Save(); err != nil {
			panic(err)
		}

		c.Response().Header.Set("Hx-Redirect", "/")
		return c.Redirect("/", fiber.StatusOK)
	}

	component := notification.Notification("Неверный логин или пароль", notification.NotificationFail)
	return tadapter.Render(c, component, fiber.StatusBadRequest)
}

func (h *AuthHandlerDeps) logout(c *fiber.Ctx) error {
	session, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}

	session.Delete("email")
	if err := session.Save(); err != nil {
		panic(err)
	}

	c.Response().Header.Set("Hx-Redirect", "/")
	return c.Redirect("/", fiber.StatusOK)
}
