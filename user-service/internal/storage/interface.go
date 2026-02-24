package storage

import "user-service/internal/models"

type UserStorageInterface interface {
	AddUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	GetAllUsers() (map[string]models.User, error)
	UpdateUser(id string, email string, password string) error
	DeleteUser(id string) error
	GetUserByEmail(email string) (*models.User, error)
}
