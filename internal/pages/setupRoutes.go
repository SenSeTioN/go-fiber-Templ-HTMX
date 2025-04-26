package pages

import (
	"log/slog"
	"sensetion/go-fiber/internal/pages/auth"
	"sensetion/go-fiber/internal/pages/home"
	"sensetion/go-fiber/internal/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SetupRoutesDeps struct {
	DBPool *pgxpool.Pool
	Logger *slog.Logger
	App    *fiber.App
	Store  *session.Store
}

func SetupRoutes(deps *SetupRoutesDeps) {
	//Repositories
	userRepo := users.NewUserRepository(deps.DBPool, deps.Logger)

	//Handlers
	home.NewHandler(deps.App)
	auth.NewAuthHandler(deps.App, &userRepo, deps.Store)
}
