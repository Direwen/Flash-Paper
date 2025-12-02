package tasks

import (
	"log"
	"os"
	"time"

	"github.com/direwen/flashpaper/internal/config"
	"github.com/direwen/flashpaper/internal/models"
)

func StartJanitor() {

	interval := os.Getenv("JANITOR_INTERVAL")
	if interval == "" {
		interval = "10"
	}

	// Parse the interval string into a duration
	duration, err := time.ParseDuration(interval + "s")
	if err != nil {
		log.Println("Failed to parse JANITOR_INTERVAL:", err)
		return
	}

	// Create a ticker that sends a signal on its channel at the specified interval
	ticker := time.NewTicker(duration)

	// Run the cleanup task concurrently in a goroutine
	go func() {
		// Infinite loop to continuously wait for ticker signals
		for {
			// Block and wait for the next signal from the ticker channel
			<-ticker.C

			// Trigger the cleanup function to delete expired snippets
			cleanExpiredSnippets()
		}
	}()

	log.Printf("The Janitor is on duty to clean every %s.", duration)
}

func cleanExpiredSnippets() {
	// Get database connection
	db := config.GetDB()

	// Delete snippets that are expired by time or have reached max views
	result := db.Where("expires_at < ? OR (max_views > 0 AND current_views >= max_views)", time.Now()).Delete(&models.Snippet{})
	if err := result.Error; err != nil {
		log.Println("Janitor failed to clean", err)
		return
	}

	// Log cleanup results
	if result.RowsAffected > 0 {
		log.Println("Janitor cleaned", result.RowsAffected, "expired snippets.")
	} else {
		log.Println("Janitor found no expired snippets.")
	}

}
