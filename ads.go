package carbot

import (
	"log"
	"time"
)

// new ads
// removed ID and QueryID from json tags because we don't need to store them in the json file
// type Ad struct {
// 	ID             int       `json:"-"`                // our ad ID
// 	QueryID        int       `json:"-"`                // our query ID
// 	Title          string    `json:"title"`            // ad title and subtitle
// 	Subtitle       string    `json:"subtitle"`         // ad title and subtitle
// 	Price          string    `json:"price"`            // ad price
// 	Other          string    `json:"other"`            // other ad information
// 	Description    string    `json:"description"`      // ad description
// 	ImageURLs      []string  `json:"image_urls"`       // ad images URL
// 	ListedTime     time.Time `json:"-"`                // time when the ad was listed
// 	DelistedTime   time.Time `json:"delisted_time"`    // time when the ad was delisted
// 	ScrapeTime     time.Time `json:"scrape_time"`      // time when the ad was scraped
// 	AdvertiserAdID string    `json:"advertiser_ad_id"` // advertiser ad ID
// }

type Ad struct {
	ID           int       `json:"-" toml:"-"`                         // our ad ID
	QueryID      int       `json:"-" toml:"-"`                         // our query ID
	Title        string    `json:"title" toml:"title"`                 // ad title and subtitle
	Subtitle     string    `json:"subtitle" toml:"subtitle"`           // ad title and subtitle
	Price        string    `json:"price" toml:"price"`                 // ad price
	Other        string    `json:"other" toml:"other"`                 // other ad information
	Description  string    `json:"description" toml:"description"`     // ad description
	ListedTime   time.Time `json:"-" toml:"-"`                         // time when the ad was listed
	DelistedTime time.Time `json:"delisted_time" toml:"delisted_time"` // time when the ad was delisted
	ScrapeTime   time.Time `json:"scrape_time" toml:"scrape_time"`     // time when the ad was scraped
	// GUI fields
	IsNew       bool      `json:"is_new" toml:"is_new"`             // is the ad new
	DeletedTime time.Time `json:"deleted_time" toml:"deleted_time"` // time when the ad was deleted
	//
	AdURL          string   `json:"ad_url" toml:"ad_url"`                     // ad URL
	ImageURLs      []string `json:"image_urls" toml:"image_urls"`             // ad images URL
	AdvertiserAdID string   `json:"advertiser_ad_id" toml:"advertiser_ad_id"` // advertiser ad ID
}

func printAds(ads []Ad) {
	for _, ad := range ads {
		log.Printf("[main] Ad: %v %v %v %v", ad.Title, ad.Price, ad.Subtitle, ad.Other)
	}
}
