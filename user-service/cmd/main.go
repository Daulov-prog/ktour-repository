package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/service"
	"user-service/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	userStorage := storage.NewUserPostgressStorage(database)
	userService := service.NewUserService(userStorage)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Post("/register", userHandler.HandleRegister)
	r.Get("/users/{id}", userHandler.HandleGetUsers)
	r.Get("/users", userHandler.HandleGetAllUsers)
	r.Delete("/users/{id}", userHandler.HandleDeleteUser)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", r)
}
