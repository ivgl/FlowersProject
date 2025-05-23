package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yourusername/plant-tracker/internal/config"
)

func NewDBConnection(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}
