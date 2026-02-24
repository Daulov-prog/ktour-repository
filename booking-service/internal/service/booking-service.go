package service

import (
	"booking-service/internal/models"
	"booking-service/internal/storage"
)

type BookingService struct {
	storage storage.BookingStorageInterface
}

func NewBookingService(storage storage.BookingStorageInterface) *BookingService {
	return &BookingService{
		storage: storage,
	}
}

func (b *BookingService) RegisterBooking(userID string, tourID string) (*models.Booking, error) {
	booking := models.NewBookingTour(userID, tourID)

	err := b.storage.AddBooking(booking)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (b *BookingService) GetBooking(id string) (*models.Booking, error) {
	booking, err := b.storage.GetBooking(id)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (b *BookingService) GetAllBookings() (map[string]models.Booking, error) {
	bookings, err := b.storage.GetAllBookings()
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (b *BookingService) DeleteBooking(id string) error {
	err := b.storage.DeleteBooking(id)
	if err != nil {
		return err
	}
	return nil
}
