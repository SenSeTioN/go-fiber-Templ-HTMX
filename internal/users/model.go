package users

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
}

type UserRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *slog.Logger
}
