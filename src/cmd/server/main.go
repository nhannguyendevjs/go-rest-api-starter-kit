package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Fatalf("Failed to load environment file (.env): %v", err)
		return
	}
	logger.Info("Environment file (.env) loaded successfully.")
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := gin.Default()

	app.Use(cors.Default())

	app.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := app.Group("/api")
	{
		v1.InitRoutes(api)
	}

	app.GET("/", func(c *gin.Context) {
		c.File("../../web/index.html")
	})

	app.NoRoute(func(c *gin.Context) {
		c.File("../../web/404.html")
	})

	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	go func() {
		logger.Infof("Server is starting on port %s...", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server on port %s: %v", port, err)
		}
	}()

	<-ctx.Done()
	logger.Info("Server has been stopped.")
}
