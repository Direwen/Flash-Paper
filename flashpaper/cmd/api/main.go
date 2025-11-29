package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/direwen/flashpaper/internal/config"
	"github.com/direwen/flashpaper/internal/handlers"
	"github.com/direwen/flashpaper/internal/middleware"
	"github.com/direwen/flashpaper/internal/services"
	"github.com/direwen/flashpaper/pkg/utils"
)

func main() {
	// Load Env Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Init Database Connection
	config.ConnectDB()
	db := config.GetDB()

	// Init Layers
	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)
	snippetService := services.NewSnippetService(db)
	snippetHandler := handlers.NewSnippetHandler(snippetService)

	// Init Gin Router
	r := gin.Default()

	{
		r.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":   "active",
				"codename": "Operation Smokescreen",
				"message":  "Systems Nominal. Ready to Burn.",
			})
		})
		r.POST("/auth/register", authHandler.Register)
		r.POST("/auth/login", authHandler.Login)
	}

	// Protected Routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			utils.SendSuccess(
				c,
				http.StatusOK,
				userID,
			)
		})
		protected.POST("/snippets", snippetHandler.Create)
	}

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port: " + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server Failed to start: ", err)
	}

}
