package main

import (
	"go/libs/db"

	"go/libs/tmdb"
)

type Scraper struct {
	client *tmdb.APIClient
	db     *db.DB
}

func NewScraper(db *db.DB) *Scraper {
	client := tmdb.NewAPIClient()
	return &Scraper{client: client, db: db}
}

func (s *Scraper) ScrapeMovie() error {
	return nil
}