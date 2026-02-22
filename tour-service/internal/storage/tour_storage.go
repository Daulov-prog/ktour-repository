package storage

import (
	"errors"
	"tour-service/internal/models"
)

type TourStorage struct {
	tourStorage map[string]models.Tour
}

func NewTourStorage() *TourStorage {
	return &TourStorage{
		tourStorage: make(map[string]models.Tour),
	}
}

func (t *TourStorage) AddTour(tour *models.Tour) error {
	if _, ok := t.tourStorage[tour.ID]; ok {
		return errors.New("Тур уже есть")
	}

	t.tourStorage[tour.ID] = *tour
	return nil
}

func (t *TourStorage) GetTourStorage(id string) (*models.Tour, error) {
	tour, ok := t.tourStorage[id]
	if !ok {
		return nil, errors.New("Такого тура нет!")
	}
	return &tour, nil
}

func (t *TourStorage) GetAllToursStorage() map[string]models.Tour {
	tmp := make(map[string]models.Tour, len(t.tourStorage))

	for k, v := range t.tourStorage {
		tmp[k] = v
	}
	return tmp
}

func (t *TourStorage) UpdateTourStatus(id string,
	description string,
	maxSlots int,
	days int,
	price float64,
	inStock bool,
) error {
	tour, ok := t.tourStorage[id]
	if !ok {
		return errors.New("Такого тура не существует")
	}

	tour.Description = description
	tour.MaxSlots = maxSlots
	tour.Days = days
	tour.Price = price
	tour.InStock = inStock
	t.tourStorage[id] = tour

	return nil
}

func (t *TourStorage) DeleteTour(id string) error {
	_, ok := t.tourStorage[id]
	if !ok {
		return errors.New("Такого туроа не существует")
	}

	delete(t.tourStorage, id)
	return nil
}
