package storage

import (
	"database/sql"
	"user-service/internal/models"
)

type UserPostgresStorage struct {
	db *sql.DB
}

func NewUserPostgressStorage(db *sql.DB) *UserPostgresStorage {
	return &UserPostgresStorage{db: db}
}

func (u *UserPostgresStorage) AddUser(user *models.User) error {
	_, err := u.db.Exec(
		"INSERT INTO users (id, full_name, age, email, password, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.FullName, user.Age, user.Email, user.Password, user.CreatedAt,
	)
	return err
}

func (u *UserPostgresStorage) GetUser(id string) (*models.User, error) {
	row := u.db.QueryRow("SELECT id, full_name, age, email, password, created_at FROM users WHERE id = $1", id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.FullName, &user.Age, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgresStorage) GetAllUsers() (map[string]models.User, error) {
	rows, err := u.db.Query("SELECT id, full_name, age, email, password, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tmp := make(map[string]models.User)
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.FullName, &user.Age, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		tmp[user.ID] = user
	}
	return tmp, nil

}

func (u *UserPostgresStorage) GetUserByEmail(email string) (*models.User, error) {
	row := u.db.QueryRow("SELECT id, full_name, age, email, password, created_at FROM users WHERE email = $1", email)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.FullName, &user.Age, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserPostgresStorage) UpdateUser(id string, email string, password string) error {
	_, err := u.db.Exec(
		"UPDATE users SET email = $1, password = $2 WHERE id = $3", email, password, id,
	)
	return err

}

func (u *UserPostgresStorage) DeleteUser(id string) error {
	_, err := u.db.Exec(
		"DELETE FROM users WHERE id = $1", id)
	return err
}
