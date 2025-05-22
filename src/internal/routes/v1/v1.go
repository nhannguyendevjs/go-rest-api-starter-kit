package v1

import (
	"github.com/gin-gonic/gin"

	"nhannguyen/gorest/internal/routes/v1/handlers"
)

func Init(appRouter *gin.RouterGroup) {
	v1Router := appRouter.Group("/v1")

	// Ping routes
	pingRoutes := v1Router.Group("/ping")
	{
		pingRoutes.GET("", handlers.GetPing)
	}

	// Book routes
	bookRoutes := v1Router.Group("/books")
	{
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.POST("", handlers.CreateBook)
	}
}
