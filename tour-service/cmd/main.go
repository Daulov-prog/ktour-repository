package main

import (
	"fmt"
	"log"
	"net/http"
	"tour-service/internal/db"
	"tour-service/internal/handler"
	"tour-service/internal/service"
	"tour-service/internal/storage"

	"github.com/go-chi/chi"
)

func main() {
	dataBase, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	tourStorage := storage.NewTourPostgressStorage(dataBase)
	tourService := service.NewTourService(tourStorage)
	tourHandler := handler.NewTourHandler(tourService)

	r := chi.NewRouter()
	r.Post("/register", tourHandler.HandleRegister)
	r.Get("/tours/{id}", tourHandler.HandleGetTour)
	r.Get("/tours", tourHandler.HandleGetAllTours)
	r.Delete("/tours/{id}", tourHandler.HandleDeleteTour)

	fmt.Println("Сервер запущен на порту 8081")
	http.ListenAndServe(":8081", r)
}
