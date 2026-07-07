package service

import (
	"errors"
	"fmt"

	"github.com/amirka7777/clock-pay/internal/repository"
	"github.com/amirka7777/clock-pay/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(input models.RegisterInput) error {

	if input.Username == "" || input.Password == "" {
		return errors.New("Имя пользователя или пароль не могут быть пустыми")
	}

	if input.Password != input.ConfirmPassword {
		return errors.New("Пароли не совпадают")
	}

	if len(input.Password) < 6 {
		return errors.New("Длина пароля должна быть не менее 6 символов")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Ошибка при хешировании пароля")
	}

	user := &models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return fmt.Errorf("Ошибка при создания пользователя в базе данных: %v", err)
	}

	return nil

}

func (s *AuthService) Login(input models.LoginInput) (*models.User, error) {

	if input.Username == "" || input.Password == "" {
		return nil, errors.New("Имя пользователя или пароль не могут быть пустыми")
	}

	user, err := s.userRepo.GetByUsername(input.Username)
	if err != nil {
		return nil, errors.New("Ошибка логина или пароля")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("Ошибка лоига или пароля")
	}

	return user, nil
}
