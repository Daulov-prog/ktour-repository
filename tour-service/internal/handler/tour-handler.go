package handler

import (
	"encoding/json"
	"net/http"
	"tour-service/internal/dto"
	"tour-service/internal/service"

	"github.com/go-chi/chi"
)

type TourHandler struct {
	service *service.TourService
}

func NewTourHandler(service *service.TourService) *TourHandler {
	return &TourHandler{
		service: service,
	}
}

func (h *TourHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterTourRequest
	json.NewDecoder(r.Body).Decode(&req)

	tour, err := h.service.RegisterTour(req.Country, req.City, req.Description, req.MaxSlots, req.Days, req.Price, req.Type, req.InStock)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := dto.TourResponse{
		ID:          tour.ID,
		Country:     tour.Country,
		City:        tour.City,
		Description: tour.Description,
		Days:        tour.Days,
		Price:       tour.Price,
		Type:        tour.Type,
		InStock:     tour.InStock,
		CreatedAt:   tour.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (h *TourHandler) HandleGetTour(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	tour, err := h.service.GetTour(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := dto.TourResponse{
		ID:          tour.ID,
		Country:     tour.Country,
		City:        tour.City,
		Description: tour.Description,
		Days:        tour.Days,
		Price:       tour.Price,
		Type:        tour.Type,
		InStock:     tour.InStock,
		CreatedAt:   tour.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TourHandler) HandleGetAllTours(w http.ResponseWriter, r *http.Request) {
	tours, err := h.service.GetAllTours()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tours)
}

func (h *TourHandler) HandleDeleteTour(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteTour(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
