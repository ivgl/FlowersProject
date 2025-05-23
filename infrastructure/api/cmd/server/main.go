package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/plant-tracker/internal/config"
	"github.com/yourusername/plant-tracker/internal/database"
	"github.com/yourusername/plant-tracker/internal/handlers"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Подключение к базе данных
	db, err := database.NewDBConnection(cfg)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Инициализация репозитория
	plantRepo := database.NewPlantRepository(db)

	// Инициализация роутера
	r := gin.Default()

	// Инициализация обработчиков
	plantHandler := handlers.NewPlantHandler(plantRepo)

	// Маршруты
	r.GET("/api/health", handlers.HealthCheck)
	r.GET("/api/plants", plantHandler.GetPlants)
	r.POST("/api/plants", plantHandler.CreatePlant)
	r.PUT("/api/plants/:plant_id", plantHandler.UpdatePlant)

	// Запуск сервера
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
