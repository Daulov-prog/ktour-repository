package storage

import "booking-service/internal/models"

type BookingStorageInterface interface {
	AddBooking(booking *models.Booking) error
	GetBooking(id string) (*models.Booking, error)
	GetAllBookings() (map[string]models.Booking, error)
	UpdateBooking(id string, userID string, status string) error
	DeleteBooking(id string) error
}
