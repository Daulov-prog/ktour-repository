package service

import (
	"tour-service/internal/models"
	"tour-service/internal/storage"
)

type TourService struct {
	storage storage.TourStorageInterface
}

func NewTourService(storage *storage.TourStorageInterface) *TourService {
	return &TourService{
		storage: *storage,
	}
}

func (t *TourService) RegisterTour(
	country string,
	city string,
	description string,
	maxSlots int,
	days int,
	price float64,
	tourType string,
	inStock bool) (*models.Tour, error) {

	tour, err := models.NewTour(
		country,
		city,
		description,
		maxSlots,
		days,
		price,
		tourType,
		inStock)
	if err != nil {
		return nil, err
	}

	err = t.storage.AddTour(tour)
	return tour, nil
}

func (t *TourService) GetTour(id string) (*models.Tour, error) {
	tour, err := t.storage.GetTour(id)
	if err != nil {
		return nil, err
	}

	return tour, nil
}

func (t *TourService) GetAllTours() (map[string]models.Tour, error) {
	tours, err := t.storage.GetAllTours()
	if err != nil {
		return nil, err
	}
	return tours, nil
}

func (t *TourService) DeleteTour(id string) error {
	err := t.storage.DeleteTour(id)
	if err != nil {
		return err
	}
	return nil
}
