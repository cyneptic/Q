package entities

import (
	"time"
)

type User struct {
	DBModel
	Name        string    `gorm:"size:255" json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `gorm:"unique" json:"phone_number"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Disabled    bool      `json:"disabled"`
}
