package v1

import (
	"nhannguyen/gorest/internal/routes/v1/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(api *gin.RouterGroup) {
	apiV1 := api.Group("/v1")

	// Ping routes
	pingRoutes := apiV1.Group("/ping")
	{
		pingRoutes.GET("", handlers.GetPing)
	}

	// Book routes
	bookRoutes := apiV1.Group("/books")
	{
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.POST("", handlers.CreateBook)
	}
}
