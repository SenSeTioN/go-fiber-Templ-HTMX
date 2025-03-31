package users

import (
	"context"
	"fmt"
	"log/slog"
	"sensetion/go-fiber/internal/types"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *slog.Logger
}

func NewUserRepository(dbpool *pgxpool.Pool, customLogger *slog.Logger) UserRepository {
	return UserRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

func (r *UserRepository) CreateUser(form types.AuthRegisterForm) error {
	query := `INSERT INTO users (id, name, email, password, created_at) VALUES (@id, @name, @email, @password, @created_at)`

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		r.CustomLogger.Error("не удалось захешировать пароль", "error", err)
		return fmt.Errorf("не удалось захешировать пароль: %w", err)
	}

	args := pgx.NamedArgs{
		"id":         uuid.New(),
		"name":       form.Name,
		"email":      form.Email,
		"password":   string(hashedPassword),
		"created_at": time.Now(),
	}

	_, err = r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (string, error) {
	var user User
	query := `SELECT email FROM users WHERE email = @email`

	args := pgx.NamedArgs{
		"email": email,
	}

	err := r.Dbpool.QueryRow(context.Background(), query, args).Scan(&user.Email)
	if err != nil {
		return "", fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	return user.Email, nil
}
