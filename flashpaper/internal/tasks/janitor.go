package tasks

import (
	"log"
	"time"

	"github.com/direwen/flashpaper/internal/config"
	"github.com/direwen/flashpaper/internal/models"
)

func StartJanitor() {

	// Create a ticker that sends a signal on its channel every 5 seconds
	ticker := time.NewTicker(5 * time.Second)

	// Run the cleanup task concurrently in a goroutine
	go func() {
		// Infinite loop to continuously wait for ticker signals
		for {
			// Block and wait for the next signal from the ticker channel
			// When the signal arrives (after 5 seconds), execution continues
			<-ticker.C

			// Trigger the cleanup function to delete expired snippets
			cleanExpiredSnippets()
		}
	}()

	log.Println("The Janitor is on duty to clean every 5 minutes.")
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
