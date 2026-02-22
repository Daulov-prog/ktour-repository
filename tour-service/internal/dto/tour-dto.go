package dto

import "time"

type RegisterTourRequest struct {
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Description string  `json:"description"`
	MaxSlots    int     `json:"max_slots"`
	Days        int     `json:"days"`
	Price       float64 `json:"price"`
	Type        string  `json:"type"`
	InStock     bool    `json:"in_stock"`
}

type TourResponse struct {
	ID          string
	Country     string
	City        string
	Description string
	Days        int
	Price       float64
	Type        string
	InStock     bool
	CreatedAt   time.Time
}
