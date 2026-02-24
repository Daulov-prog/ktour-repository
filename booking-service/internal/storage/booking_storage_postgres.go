package storage

import (
	"booking-service/internal/models"
	"database/sql"
)

type BookingStorage struct {
	db *sql.DB
}

func NewBookingStorage(db *sql.DB) *BookingStorage {
	return &BookingStorage{db: db}
}

func (b *BookingStorage) AddBooking(booking *models.Booking) error {
	_, err := b.db.Exec(
		"INSERT INTO bookings (id, user_id, tour_id, status, created_at) VALUES ($1, $2, $3, $4, $5)",
		booking.ID, booking.UserID, booking.TourID, booking.Status, booking.CreatedAt,
	)
	return err
}

func (b *BookingStorage) GetBooking(id string) (*models.Booking, error) {
	row := b.db.QueryRow("SELECT id, user_id, tour_id, status FROM bookings WHERE id = $1", id)

	booking := &models.Booking{}
	err := row.Scan(&booking.ID, &booking.UserID, &booking.TourID, &booking.Status)
	if err != nil {
		return nil, err
	}

	return booking, err
}

func (b *BookingStorage) GetAllBookings() (map[string]models.Booking, error) {
	rows, err := b.db.Query("SELECT id, user_id, tour_id, status FROM bookings")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tmp := make(map[string]models.Booking)
	for rows.Next() {
		booking := models.Booking{}

		err := rows.Scan(&booking.ID, &booking.UserID, &booking.TourID, &booking.Status)
		if err != nil {
			return nil, err
		}
		tmp[booking.ID] = booking
	}
	return tmp, nil
}

func (b *BookingStorage) UpdateBooking(id string, userID string, status string) error {
	_, err := b.db.Exec(
		"UPDATE bookings SET user_id = $2, status = $3 WHERE id = $1", id, userID, status,
	)
	return err
}

func (b *BookingStorage) DeleteBooking(id string) error {
	_, err := b.db.Exec(
		"DELETE FROM bookings WHERE id = $1", id,
	)
	return err
}
