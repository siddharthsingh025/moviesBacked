package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	ID          string    `gorm:"primaryKey"`
	ReleaseDate string    `json:"release_date" validate:"required"`
	Title       string    `json:"title"`
	Director    Director  `json:"director" gorm:"foreignKey:ID"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
