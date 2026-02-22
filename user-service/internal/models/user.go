package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	FullName  string
	Age       int
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(
	fullName string,
	age int,
	email string,
	password string,
) (*User, error) {
	id := uuid.New().String()

	if fullName == "" {
		return nil, errors.New("Пустое поле")
	}

	if age < 0 || age > 150 {
		return nil, errors.New("Невозможное значение")
	}

	if email == "" {
		return nil, errors.New("Пустое поле")
	}

	if password == "" {
		return nil, errors.New("Пустое поле")
	}

	return &User{
		ID:       id,
		FullName: fullName,
		Age:      age,
		Email:    email,
		Password: password,

		CreatedAt: time.Now(),
	}, nil
}
