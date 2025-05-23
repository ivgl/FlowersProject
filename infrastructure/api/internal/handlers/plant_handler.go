package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/plant-tracker/internal/database"
	"github.com/yourusername/plant-tracker/internal/models"
)

type PlantHandler struct {
	repo *database.PlantRepository
}

func NewPlantHandler(repo *database.PlantRepository) *PlantHandler {
	return &PlantHandler{repo: repo}
}

func (h *PlantHandler) GetPlants(c *gin.Context) {
	plants, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plants)
}

func (h *PlantHandler) CreatePlant(c *gin.Context) {
	var req models.CreatePlantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plant, err := h.repo.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plant)
}

func (h *PlantHandler) UpdatePlant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("plant_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plant ID"})
		return
	}

	var req models.UpdatePlantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plant, err := h.repo.Update(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plant)
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
