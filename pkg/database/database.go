package database

import (
	"context"
	"log/slog"
	"sensetion/go-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool(config *config.DatabaseConfig, logger *slog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error("Не удалось подключиться к базе данных", "error", err)
		panic(err)
	}

	logger.Info("Подключение к базе данных установлено")

	return dbpool
}
