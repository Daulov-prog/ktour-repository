package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        string
	UserID    string
	TourID    string
	Status    string
	CreatedAt time.Time
}

func NewBookingTour(userId string, tourId string) *Booking {
	return &Booking{
		ID:        uuid.New().String(),
		UserID:    userId,
		TourID:    tourId,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
}
