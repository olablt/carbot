package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olablt/carbot"
)

var confDir = "./apps/06-ads-db/conf/"

func main() {
	setupPath(confDir)

	ap := carbot.NewAutoplius()

	reg := carbot.NewAdvertisersRegistry()
	reg.Add(ap)

	b := carbot.NewBrowser()

	for {
		// process queries
		for _, q := range queries {
			scrapedAds, err := reg.ProcessQuery(q, b)
			if err != nil {
				log.Println("[main] Error processing query:", err)
			} else {
				// log summary
				log.Printf("[main] Scraped Ads: %v", len(scrapedAds))
			}

			if len(scrapedAds) > 0 {
				// update ads cache
				updateAdsCache(q, scrapedAds)
			} else {
				// log.Println("[main] No ads found")
			}

		}
		// wait and repeat all queries
		// time.Sleep(5 * time.Second)
		break
	}
	b.CloseBrowser()
}
