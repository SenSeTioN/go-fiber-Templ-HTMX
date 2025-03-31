package pages

import (
	"log/slog"
	"sensetion/go-fiber/internal/pages/auth"
	"sensetion/go-fiber/internal/pages/home"
	"sensetion/go-fiber/internal/users"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SetupRoutesDeps struct {
	DBPool *pgxpool.Pool
	Logger *slog.Logger
	App    *fiber.App
}

func SetupRoutes(args *SetupRoutesDeps) {
	//Repositories
	userRepo := users.NewUserRepository(args.DBPool, args.Logger)

	//Handlers
	home.NewHandler(args.App)
	auth.NewAuthHandler(args.App, &userRepo)
}
