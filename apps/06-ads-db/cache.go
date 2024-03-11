package main

import (
	"log"
	"time"

	"github.com/olablt/carbot"
)

// load ads from file cache and compare with new ads
func updateAdsCache(q *carbot.Query, scrapedAds []carbot.Ad) {
	carDB, err := carbot.OpenCarbotDB(confDir)
	if err != nil {
		log.Fatal("[DEBUG] error opening database", err)
	}

	// get all ads from the database
	cachedAds, err := carDB.GetAdsByQueryID(q.ID)
	if err != nil {
		log.Fatal("[DEBUG] error getting tasks", err)
	}
	log.Printf("[cache] Loaded cached ads:%v for query:%v", len(cachedAds), q.ID)

	// loop through scrapedAds and compare with cached ads
	// if there are differences, add to cachedAds and write to json file
	newAds := []carbot.Ad{}
	for _, ad := range scrapedAds {
		found := false
		for _, ad2 := range cachedAds {
			if ad.AdvertiserAdID == ad2.AdvertiserAdID {
				// if ad.ImageURLs[0] == ad2.ImageURLs[0] {
				found = true
				break
			}
		}
		if !found {
			ad.QueryID = q.ID
			newAds = append(newAds, ad)
		}
	}

	// check for delisted ads
	// delistedCount := 0
	delistedAds := []carbot.Ad{}
	for _, ad := range cachedAds {
		found := false
		for _, ad2 := range scrapedAds {
			if ad.AdvertiserAdID == ad2.AdvertiserAdID {
				found = true
				break
			}
		}
		// if cached ad is not found in scraped ads, mark it as delisted
		if !found {
			if ad.DelistedTime.IsZero() {
				ad.DelistedTime = time.Now()
				delistedAds = append(delistedAds, ad)
			}
		}
	}

	if len(newAds) > 0 || len(delistedAds) > 0 {
		log.Printf("[cache] found %v new ads", len(newAds))
		log.Printf("[cache] found %v delisted ads", len(delistedAds))
		if len(newAds) > 0 {
			// insert new ads into the database
			log.Println("[cache] inserting new ads into the database")
			carDB.InsertMany(newAds)
			// for _, ad := range newAds {
			// 	id, err := carDB.InsertOne(&ad)
			// 	if err != nil {
			// 		log.Fatal("[DEBUG] error inserting ad", err)
			// 	}
			// 	log.Println("ID:", id)
			// 	ad.ID = id
			// }
		}
		if len(delistedAds) > 0 {
			// update delisted ads in the database
			// carDB.UpdateMany(delistedAds)
			// update one by one
			for _, ad := range delistedAds {
				log.Println("[cache] updating delisted ad in the database", ad.ID, ad.DelistedTime)
				err := carDB.Update(ad)
				if err != nil {
					log.Fatal("[DEBUG] error updating ad", err)
				}
			}
		}
	} else {
		log.Println("[cache] no new or delisted ads found")
	}
}
