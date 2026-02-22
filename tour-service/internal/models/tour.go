package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Tour struct {
	ID          string
	Country     string
	City        string
	Description string
	MaxSlots    int
	Days        int
	Price       float64
	Type        string
	InStock     bool
	CreatedAt   time.Time
}

func NewTour(
	country string,
	city string,
	description string,
	maxSlots int,
	days int,
	price float64,
	tourType string,
	inStock bool,
) (*Tour, error) {
	id := uuid.New().String()

	if country == "" {
		return nil, errors.New("Поле не может быть пустым")
	}

	if city == "" {
		return nil, errors.New("Поле не может быть пустым")
	}

	if description == "" {
		return nil, errors.New("Поле не может быть пустым")
	}

	if maxSlots <= 0 {
		return nil, errors.New("Поле не может быть пустым")
	}

	if days <= 0 {
		return nil, errors.New("Невозможное значение")
	}

	if price <= 0 {
		return nil, errors.New("Некорректная цена")
	}

	if tourType == "" {
		return nil, errors.New("Поле не может быть пустым")
	}

	return &Tour{
		ID:          id,
		Country:     country,
		City:        city,
		Description: description,
		MaxSlots:    maxSlots,
		Days:        days,
		Price:       price,
		Type:        tourType,
		InStock:     inStock,
		CreatedAt:   time.Now(),
	}, nil
}
