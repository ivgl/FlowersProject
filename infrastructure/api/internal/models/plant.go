package models

import (
	"time"
)

type Plant struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	Species           string    `json:"species"`
	WateringFrequency int       `json:"watering_frequency"`
	LastWatered       time.Time `json:"last_watered"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreatePlantRequest struct {
	Name              string    `json:"name" binding:"required"`
	Species           string    `json:"species" binding:"required"`
	WateringFrequency int       `json:"watering_frequency" binding:"required"`
	LastWatered       time.Time `json:"last_watered" binding:"required"`
}

type UpdatePlantRequest struct {
	Name              string    `json:"name" binding:"required"`
	Species           string    `json:"species" binding:"required"`
	WateringFrequency int       `json:"watering_frequency" binding:"required"`
	LastWatered       time.Time `json:"last_watered" binding:"required"`
}
