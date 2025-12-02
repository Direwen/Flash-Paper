package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/direwen/flashpaper/internal/config"
	"github.com/direwen/flashpaper/internal/handlers"
	"github.com/direwen/flashpaper/internal/middleware"
	"github.com/direwen/flashpaper/internal/services"
	"github.com/direwen/flashpaper/internal/tasks"
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

	// Health Check: Database Connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}
	log.Println("Database connection verified.")

	// Start Background Task
	tasks.StartJanitor()

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
		r.GET("/snippets/:id", snippetHandler.Get)
	}

	// Protected Routes
	protected := r.Group("")
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
		protected.DELETE("/snippets/:id", snippetHandler.Delete)
	}

	// Get port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port: " + port)

	// if err := r.Run(":" + port); err != nil {
	// 	log.Fatal("Server Failed to start: ", err)
	// }

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Server starts in background goroutine so main goroutine can continue
	go func() {
		// server.ListenAndServe runs forever (infinite loop) until shutdown
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Failed to start: %v", err)
		}
	}()

	// Create a channel to hold one os.Signal value
	quit := make(chan os.Signal, 1)
	// When the OS sends SIGINT (Ctrl+C) or SIGTERM (docker stop) to this process,
	// signal.Notify catches it and sends that signal into the quit channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Block main goroutine here until a signal is received from the channel
	<-quit
	log.Println("Shutting down server...")

	// Create a context that automatically expires after 10 seconds (sets a deadline)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ensures cancel() is called when main() exits to release context resources
	defer cancel()

	// Graceful shutdown: stops accepting new requests, waits for active requests
	// to finish (up to 10 seconds), then closes all connections
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed:", err)
	}

	// Close the database connection
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Failed to close database connection: ", err)
	}

	log.Println("Server exited successfully")

}
