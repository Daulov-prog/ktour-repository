package models

import "time"

type Booking struct {
	BookingID     string
	UserID        string
	TourID        string
	BookingTime   time.Time
	BookingStatus string
}

// package models

// import "time"

// type Tour struct {
// 	ID          string
// 	Country     string
// 	City        string
// 	Description string
// 	MaxSlots    int
// 	Days        int
// 	Price       float64
// 	Type        string
// 	InStock     bool
// 	CreatedAt   time.Time
// }

// package models

// import "time"

// type User struct {
// 	ID        string
// 	FullName  string
// 	Age       int
// 	Email     string
// 	Password  string
// 	CreatedAt time.Time
// }
