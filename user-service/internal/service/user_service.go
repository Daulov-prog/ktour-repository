package service

import (
	"user-service/internal/models"
	"user-service/internal/storage"
)

type UserService struct {
	storage storage.UserStorageInterface
}

func NewUserService(storage storage.UserStorageInterface) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (u *UserService) RegisterUser(fullName string, age int, email string, password string) (*models.User, error) {
	user, err := models.NewUser(fullName, age, email, password)
	if err != nil {
		return nil, err
	}

	err = u.storage.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (u *UserService) GetUser(id string) (*models.User, error) {
	user, err := u.storage.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetAllUsers() (map[string]models.User, error) {
	users, err := u.storage.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, err
}

func (u *UserService) DeleteUser(id string) error {
	err := u.storage.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
