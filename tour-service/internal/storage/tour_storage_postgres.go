package storage

import (
	"database/sql"
	"tour-service/internal/models"
)

type TourPostgressStorage struct {
	db *sql.DB
}

func NewTourPostgressStorage(db *sql.DB) *TourPostgressStorage {
	return &TourPostgressStorage{db: db}
}

func (t *TourPostgressStorage) AddTour(tour *models.Tour) error {
	_, err := t.db.Exec(
		"INSERT INTO tours (id, country, city, description, max_slots, days, price, type, in_stock, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		tour.ID, tour.Country, tour.City, tour.Description, tour.MaxSlots, tour.Days, tour.Price, tour.Type, tour.InStock, tour.CreatedAt,
	)
	return err
}

func (t *TourPostgressStorage) GetTour(id string) (*models.Tour, error) {
	row := t.db.QueryRow("SELECT id, country, city,  description, max_slots, days, price, type, in_stock, created_at FROM tours WHERE id = $1", id)

	tour := &models.Tour{}
	err := row.Scan(&tour.ID, &tour.Country, &tour.City, &tour.Description, &tour.MaxSlots, &tour.Days, &tour.Price, &tour.Type, &tour.InStock, &tour.CreatedAt)
	if err != nil {
		return nil, err
	}

	return tour, nil
}

func (t *TourPostgressStorage) GetAllTours() (map[string]models.Tour, error) {
	rows, err := t.db.Query("SELECT id, country, city, description, max_slots, days, price, type, in_stock, created_at FROM tours")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tmp := make(map[string]models.Tour)
	for rows.Next() {
		tour := models.Tour{}

		err := rows.Scan(&tour.ID, &tour.Country, &tour.City, &tour.Description, &tour.MaxSlots, &tour.Days, &tour.Price, &tour.Type, &tour.InStock, &tour.CreatedAt)
		if err != nil {
			return nil, err
		}
		tmp[tour.ID] = tour
	}
	return tmp, nil
}

func (t *TourPostgressStorage) UpdateTour(id string, description string, maxSlots int, days int, price float64, inStock bool) error {
	_, err := t.db.Exec(
		"UPDATE tours SET description = $1, max_slots = $2, days = $3, price = $4, in_stock =$5 WHERE id = $6", description, maxSlots, days, price, inStock, id,
	)
	return err
}

func (t *TourPostgressStorage) DeleteTour(id string) error {
	_, err := t.db.Exec(
		"DELETE FROM tours WHERE id = $1", id,
	)
	return err
}
