package main

// // INSERT
// newAd := &carbot.Ad{}
// newAd.Title = "new title"
// newAd.ImageURLs = []string{"new image3"}
// newAd.QueryID = 3
// id, err := carDB.Insert(newAd)
// if err != nil {
// 	log.Fatal("[DEBUG] error inserting task", err)
// }
// log.Println("ID:", id)
// newAd.ID = id

// // UPDATE
// newAd := &carbot.Ad{}
// newAd.ID = 7
// // newAd.QueryID = 4
// newAd.Title = "new title3"
// newAd.ImageURLs = []string{"new image5"}
// newAd.ScrapeTime = time.Now()
// err = carDB.Update(*newAd)
// if err != nil {
// 	log.Fatal("[DEBUG] error updating task", err)
// }

// // GET
// id := 7
// ad, err := carDB.GetAd(id)
// if err != nil {
// 	log.Fatal("[DEBUG] error getting task", err)
// }
// log.Printf("Ad: %+v", ad)

// // DELETE BY QUERY ID
// queryID := 0
// err = carDB.DeleteByQueryID(queryID)
// if err != nil {
// 	log.Fatal("[DEBUG] error deleting tasks", err)
// }

// // GET ALL
// ads, err := carDB.GetAds()
// if err != nil {
// 	log.Fatal("[DEBUG] error getting tasks", err)
// }
// log.Println("GET ALL")
// for _, ad := range ads {
// 	log.Printf("Ad: #%v %v %v %v", ad.ID, ad.QueryID, ad.Title, ad.Subtitle)
// }

// // GET BY QUERY ID
// queryID := 3
// ads, err = carDB.GetAdsByQueryID(queryID)
// if err != nil {
// 	log.Fatal("[DEBUG] error getting tasks", err)
// }
// log.Println("GET BY QUERY ID")
// for _, ad := range ads {
// 	log.Printf("Ad: #%v %v %v %v", ad.ID, ad.QueryID, ad.Title, ad.Subtitle)
// }
