package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/amirka7777/clock-pay/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {

	query := `INSERT INTO users (username, password_hash) VALUES (?, ?)`
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("Ошибка при создании пользователя: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {

	query := `SELECT id, username, password_hash, created_at FROM users WHERE username = ?`

	row := r.db.QueryRow(query, username)
	var user models.User

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Пользователь с таким ником не найден")
		}
		return nil, fmt.Errorf("Ошибка при сканировании данных пользователя: %w", err)
	}

	return &user, nil

}
