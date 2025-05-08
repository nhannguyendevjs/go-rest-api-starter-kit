package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nhannguyen/gorest/internal/routes/v1/models"
)

// GetPing godoc
// @Summary Get ping
// @Description Get ping
// @Tags Utilities
// @Produce json
// @Success 200 {object} models.Ping
// @Router /api/v1/ping [get]
func GetPing(c *gin.Context) {
	pong := models.Ping{
		Success: true,
		Data:    "pong",
	}
	c.JSON(http.StatusOK, pong)
}
