package handler

import (
	"encoding/json"
	"net/http"

	"github.com/amirka7777/clock-pay/internal/service"
	"github.com/amirka7777/clock-pay/models"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var input models.RegisterInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errors": "невалидное тело запроса"})
		return
	}

	err = h.authService.Register(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errors": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "регистрация прошла успешно"})

}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var input models.LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errors": "невалидное тело запроса"})
		return
	}

	user, err := h.authService.Login(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"errors": err.Error()})
		return
	}

	// временное решение, в дальнейшей - jwt-токен
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "авторизация прошла успешно",
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
		},
	})

}
