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
	"nhannguyen/gorest/internal/handlers"
	_ "nhannguyen/gorest/pkg/logger"
)

// @title Book API
// @version 1.0
// @description This is a sample server for managing books.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load env variables
	loadEnv()
	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")
	// System signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Initial gin app
	app := initApp(ginMode)
	// Start HTTP server
	go startServer(port, app)
	// Clean up
	<-ctx.Done()
	cleanup()
}

func loadEnv() {
	err := godotenv.Load("../../.env")

	if err != nil {
		logger.Fatalf("Failed to load environment file (.env): %v", err)
		return
	}

	logger.Info("Environment file (.env) loaded successfully.")
}

func initApp(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	app := gin.Default()

	// Use middlewares
	app.Use(cors.Default())

	// docs
	app.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// api
	appRouter := app.Group("/api/v1")
	{
		// Ping routes
		pingRoutes := appRouter.Group("/ping")
		{
			pingRoutes.GET("", handlers.GetPing)
		}
		// Book routes
		bookRoutes := appRouter.Group("/books")
		{
			bookRoutes.GET("", handlers.GetBooks)
			bookRoutes.POST("", handlers.CreateBook)
		}
	}

	// root
	app.GET("/", func(context *gin.Context) {
		context.File("../../web/index.html")
	})

	// not-found
	app.NoRoute(func(context *gin.Context) {
		context.File("../../web/404.html")
	})

	return app
}

func startServer(port string, app *gin.Engine) {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	logger.Infof("Server is starting on port %s...", port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Failed to start server on port %s: %v", port, err)
	}
}

func cleanup() {
	logger.Info("Server has been stopped.")
}
