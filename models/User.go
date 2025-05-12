package model

import (
	"time"
)

type User struct {
	ID        uint      `json:"id"`                          // Standard field for the primary key
	Name      string    `json:"name" binding:"required"`     // A regular string field
	Email     string    `json:"email" binding:"required"`    // A pointer to a string, allowing for null values
	Phone     string    `json:"phone" binding:"required"`    // A pointer to a string, allowing for null values
	Password  string    `json:"password" binding:"required"` // A pointer to a string, allowing for null values
	CreatedAt time.Time `json:"created_at"`                  // Automatically managed by GORM for creation time
	UpdatedAt time.Time `json:"updated_at"`                  // Automatically managed by GORM for update time
}
