package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/olablt/carbot"
)

// load ads from file cache and compare with new ads
func updateAdsCache(cacheFilename string, scrapedAds []carbot.Ad) {
	// unmarshal allAnnouncements from json file
	var cachedAds []carbot.Ad
	jsonFile, err := os.ReadFile(cacheFilename)
	if err == nil {
		json.Unmarshal(jsonFile, &cachedAds)

		// loop through scrapedAds and compare with cached ads
		// if there are differences, add to cachedAds and write to json file
		newCount := 0
		for _, ad := range scrapedAds {
			found := false
			for _, ad2 := range cachedAds {
				if ad.ImageURLs[0] == ad2.ImageURLs[0] {
					found = true
					break
				}
			}
			if !found {
				newCount++
				cachedAds = append(cachedAds, ad)
			}
		}

		// check for delisted ads
		delistedCount := 0
		for i, ad := range cachedAds {
			found := false
			for _, ad2 := range scrapedAds {
				if ad.ImageURLs[0] == ad2.ImageURLs[0] {
					found = true
					break
				}
			}
			// if cached ad is not found in scraped ads, mark it as delisted
			if !found {
				delistedCount++
				cachedAds[i].DelistedTime = time.Now()
			}
		}

		if newCount > 0 || delistedCount > 0 {
			log.Printf("[cache] found %v new ads", newCount)
			log.Printf("[cache] found %v delisted ads", delistedCount)
			// marshal cachedAds to json file
			jsonads2, _ := json.MarshalIndent(cachedAds, "", "  ")
			os.WriteFile(cacheFilename, jsonads2, 0644)
		} else {
			log.Println("[cache] no new or delisted ads found")
		}
	} else {
		// ads cache not found, create it
		log.Println("[cache] ads cache not found, creating it")
		// check if cache file is json or toml
		if len(cacheFilename) > 5 && cacheFilename[len(cacheFilename)-5:] == ".json" {
			// log.Println("[cache] cache file is json")
			jsonads2, _ := json.MarshalIndent(scrapedAds, "", "  ")
			os.WriteFile(cacheFilename, jsonads2, 0644)
		} else {
			// log.Println("[cache] cache file is toml")
			// // toml
			// f, err := os.Create(cacheFilename)
			// if err != nil {
			// 	panic(err)
			// }
			// defer f.Close()
			// encoder := toml.NewEncoder(f)
			// err = encoder.Encode(scrapedAds)
		}
	}
}
