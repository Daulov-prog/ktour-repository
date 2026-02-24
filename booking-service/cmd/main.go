package main

import (
	"booking-service/internal/db"
	"booking-service/internal/handler"
	"booking-service/internal/service"
	"booking-service/internal/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	dataBase, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	bookingStorage := storage.NewBookingStorage(dataBase)
	bookingService := service.NewBookingService(bookingStorage)
	bookingHanlder := handler.NewTourHandler(bookingService)

	r := chi.NewRouter()
	r.Post("/bookings", bookingHanlder.HandleRegister)
	r.Get("/bookings/{id}", bookingHanlder.HandlGetBooking)
	r.Get("/bookings", bookingHanlder.HandleGetAllBooking)
	r.Delete("/bookings/{id}", bookingHanlder.HandleDeleteBooking)

	fmt.Println("Сервер запущен на порту 8082")
	http.ListenAndServe(":8082", r)
}
