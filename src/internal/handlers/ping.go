package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nhannguyen/gorest/internal/models"
)

// GetPing godoc
// @Summary Get ping
// @Description Get ping
// @Tags Utilities
// @Security BearerAuth
// @Accept  json
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /api/v1/ping [get]
func GetPing(c *gin.Context) {
	data := models.PingResponse{
		Success: true,
		Data:    "pong",
	}
	c.JSON(http.StatusOK, data)
}
