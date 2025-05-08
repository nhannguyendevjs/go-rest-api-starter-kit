package v1

import (
	"github.com/gin-gonic/gin"

	"nhannguyen/gorest/internal/routes/v1/handlers"
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
