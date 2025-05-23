package database

import (
	"database/sql"
	"time"

	"github.com/yourusername/plant-tracker/internal/models"
)

type PlantRepository struct {
	db *sql.DB
}

func NewPlantRepository(db *sql.DB) *PlantRepository {
	return &PlantRepository{db: db}
}

func (r *PlantRepository) GetAll() ([]models.Plant, error) {
	query := `
		SELECT id, name, species, watering_frequency, last_watered, created_at, updated_at 
		FROM plants 
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []models.Plant
	for rows.Next() {
		var p models.Plant
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Species,
			&p.WateringFrequency,
			&p.LastWatered,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		plants = append(plants, p)
	}

	return plants, nil
}

func (r *PlantRepository) Create(plant models.CreatePlantRequest) (models.Plant, error) {
	query := `
		INSERT INTO plants (name, species, watering_frequency, last_watered, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, name, species, watering_frequency, last_watered, created_at, updated_at`

	now := time.Now()
	var p models.Plant
	err := r.db.QueryRow(
		query,
		plant.Name,
		plant.Species,
		plant.WateringFrequency,
		plant.LastWatered,
		now,
		now,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Species,
		&p.WateringFrequency,
		&p.LastWatered,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *PlantRepository) Update(id int64, plant models.UpdatePlantRequest) (models.Plant, error) {
	query := `
		UPDATE plants 
		SET name = $1, species = $2, watering_frequency = $3, last_watered = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, name, species, watering_frequency, last_watered, created_at, updated_at`

	var p models.Plant
	err := r.db.QueryRow(
		query,
		plant.Name,
		plant.Species,
		plant.WateringFrequency,
		plant.LastWatered,
		time.Now(),
		id,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Species,
		&p.WateringFrequency,
		&p.LastWatered,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}
