package model

import "time"

type Todo struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `gorm:"not null" json:"-"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	Done      bool      `gorm:"not null" json:"done"`
}
