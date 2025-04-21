package main

import (
	_ "nhannguyen/gorest/docs" // swag docs
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	apiV1 "nhannguyen/gorest/internal/routes/v1"
	"nhannguyen/gorest/pkg/logger"
)

// @title Book API
// @version 1.0
// @description This is a sample server for managing books.
// @host localhost:8080
// @BasePath /

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Info("Error loading .env file")
	}
	logger.Log.Info("Loaded .env file")
}

func main() {
	router := gin.Default()

	// Configure CORS to only allow specific non-localhost origins
	config := cors.Config{
		AllowOrigins:     []string{"https://your-production-domain.com"}, // Only allow this
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	router.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		apiV1.InitRoutes(api)
	}

	router.GET("/", func(c *gin.Context) {
		c.File("web/index.html")
	})

	port := os.Getenv("PORT")
	router.Run(":" + port)
	logger.Log.Infof("Server started on port %s", port)
}
