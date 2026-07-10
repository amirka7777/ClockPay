package main

import (
	"log"
	"net/http"

	"github.com/amirka7777/clock-pay/internal/handler"
	"github.com/amirka7777/clock-pay/internal/repository"
	"github.com/amirka7777/clock-pay/internal/service"
)

func main() {

	db, err := repository.InitDB("./subscriptions.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth/register", authHandler.RegisterHandler)
	mux.HandleFunc("/api/auth/login", authHandler.LoginHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("сервер запустился с ошибкой: %v", err)
	}

}
