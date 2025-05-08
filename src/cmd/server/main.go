package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	logger "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "nhannguyen/gorest/docs"
	v1 "nhannguyen/gorest/internal/routes/v1"
	_ "nhannguyen/gorest/pkg/logger"
)

// @title Book API
// @version 1.0
// @description This is a sample server for managing books.
// @host localhost:8080
// @BasePath /

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Info("Error loading .env file")
		return
	}
	logger.Info("Loaded .env file")
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
		v1.InitRoutes(api)
	}

	router.GET("/", func(c *gin.Context) {
		c.File("web/index.html")
	})

	router.NoRoute(func(c *gin.Context) {
		c.File("web/404.html")
	})

	port := os.Getenv("PORT")
	router.Run(":" + port)
	logger.Infof("Server started on port %s", port)
}
