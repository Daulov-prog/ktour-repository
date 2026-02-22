package storage

import (
	"errors"
	"user-service/internal/models"
)

type UserStorage struct {
	userStorage map[string]models.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		userStorage: make(map[string]models.User),
	}
}

func (u *UserStorage) AddUser(user *models.User) error {
	if _, ok := u.userStorage[user.ID]; ok {
		return errors.New("Пользователь уже есть")
	}

	u.userStorage[user.ID] = *user
	return nil
}

func (u *UserStorage) GetUserStorage(id string) (*models.User, error) {
	user, ok := u.userStorage[id]
	if !ok {
		return nil, errors.New("Такого пользователя нет")
	}

	return &user, nil
}

func (u *UserStorage) GetAllUsers() map[string]models.User {
	tmp := make(map[string]models.User, len(u.userStorage))

	for k, v := range u.userStorage {
		tmp[k] = v
	}
	return tmp
}

func (u *UserStorage) UpdateUserStatus(id string, email string, password string) error {
	user, ok := u.userStorage[id]
	if !ok {
		return errors.New("Такого пользователя не существует")
	}

	user.Email = email
	user.Password = password
	u.userStorage[id] = user

	return nil

}

func (u *UserStorage) DeleteUser(id string) error {
	_, ok := u.userStorage[id]
	if !ok {
		return errors.New("Такого пользователя не существует")
	}

	delete(u.userStorage, id)

	return nil
}
