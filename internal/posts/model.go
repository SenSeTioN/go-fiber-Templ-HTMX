package posts

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *slog.Logger
}
