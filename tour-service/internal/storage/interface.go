package storage

import "tour-service/internal/models"

type TourStorageInterface interface {
	AddTour(tour *models.Tour) error
	GetTour(id string) (*models.Tour, error)
	GetAllTours() (map[string]models.Tour, error)
	UpdateTour(id string, description string, maxSlots int, days int, price float64, inStock bool) error
	DeleteTour(id string) error
}
