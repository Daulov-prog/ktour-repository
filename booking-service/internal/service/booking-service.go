package service

import (
	"booking-service/internal/grpc"
	"booking-service/internal/models"
	"booking-service/internal/storage"
	"errors"
)

type BookingService struct {
	storage    storage.BookingStorageInterface
	userClient *grpc.UserClient
}

func NewBookingService(storage storage.BookingStorageInterface, userClient *grpc.UserClient) *BookingService {
	return &BookingService{
		storage:    storage,
		userClient: userClient,
	}
}

func (b *BookingService) RegisterBooking(userID string, tourID string) (*models.Booking, error) {
	_, err := b.userClient.GetUser(userID)
	if err != nil {
		return nil, errors.New("Пользователь не найден")
	}

	booking := models.NewBookingTour(userID, tourID)

	err = b.storage.AddBooking(booking)
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
