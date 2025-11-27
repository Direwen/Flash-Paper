package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/direwen/flashpaper/internal/config"
)

func main() {
	// Load Env Variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Init Database Connection
	config.ConnectDB()

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
