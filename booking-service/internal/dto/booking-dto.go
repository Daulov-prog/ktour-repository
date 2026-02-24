package dto

import "time"

type RegisterBookingRequest struct {
	UserID string `json:"user_id"`
	TourID string `json:"tour_id"`
	Status string `json:"status"`
}

type BookingResponse struct {
	ID        string
	UserID    string
	TourID    string
	Status    string
	CreatedAt time.Time
}
