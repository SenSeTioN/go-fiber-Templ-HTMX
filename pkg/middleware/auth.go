package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func AuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		userMail := ""
		if email, ok := session.Get("email").(string); ok {
			userMail = email
		}
		c.Locals("email", userMail)

		userName := ""
		if name, ok := session.Get("name").(string); ok {
			userName = name
		}
		c.Locals("name", userName)

		return c.Next()
	}
}
