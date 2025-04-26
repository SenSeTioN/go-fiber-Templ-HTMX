package posts

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostRepository(dbpool *pgxpool.Pool, customLogger *slog.Logger) PostRepository {
	return PostRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *PostRepository) LastFivePosts() {}
