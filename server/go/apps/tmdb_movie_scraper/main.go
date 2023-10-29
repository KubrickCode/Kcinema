package main

import (
	"go/libs/db"
	"log"
	"time"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	scraper := NewScraper(db)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		err := scraper.ScrapeMovie()
		if err != nil {
			log.Fatalf("Task %s failed: %v", "scrap", err)
		}

		<-ticker.C
	}
}