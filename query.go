package carbot

import "time"

// add json tags to the struct fields
type Query struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	AdvertiserName  string        `json:"advertiser_name"`
	FirstURL        string        `json:"first_url"`
	LastScrapeTime  time.Time     `json:"last_scrape_time"`
	LastScrapePages int           `json:"last_scrape_pages"`
	LastScrapeAds   int           `json:"last_scrape_ads"`
	AdCache         []AdFootprint `json:"-"`
}

// ad cache is used to store all the ads that we have scraped
type AdFootprint struct {
	ID             int    // our ad ID
	QueryID        string // our query ID
	AdvertiserAdID string // advertiser ad ID
}
