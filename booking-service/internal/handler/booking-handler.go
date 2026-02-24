package handler

import (
	"booking-service/internal/dto"
	"booking-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type BookingHandler struct {
	service *service.BookingService
}

func NewTourHandler(service *service.BookingService) *BookingHandler {
	return &BookingHandler{
		service: service,
	}
}

func (h *BookingHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterBookingRequest
	json.NewDecoder(r.Body).Decode(&req)

	booking, err := h.service.RegisterBooking(req.UserID, req.TourID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := dto.BookingResponse{
		ID:        booking.ID,
		UserID:    booking.UserID,
		TourID:    booking.TourID,
		Status:    booking.Status,
		CreatedAt: booking.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *BookingHandler) HandlGetBooking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	booking, err := h.service.GetBooking(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := dto.BookingResponse{
		ID:        booking.ID,
		UserID:    booking.UserID,
		TourID:    booking.TourID,
		Status:    booking.Status,
		CreatedAt: booking.CreatedAt,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *BookingHandler) HandleGetAllBooking(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) HandleDeleteBooking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteBooking(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
