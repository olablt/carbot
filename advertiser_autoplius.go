package carbot

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Atoplius implements the AdvertiserInterface
type Autoplius struct {
	Name    string // unique name
	FormURL string // URL to the form
}

// NewAutoplius creates a new Autoplius instance
func NewAutoplius() *Autoplius {
	return &Autoplius{
		Name:    "autoplius.lt",
		FormURL: "https://autoplius.lt/paieska/naudoti-automobiliai",
	}
}

// for future debugging
func (a *Autoplius) ProcessQueryHTML(q *Query, html string) ([]Ad, error) {
	doc, _ := GoQueryFromString(html)
	return a.parseAnnouncements(doc)
}

// ProcessQuery processes the query
func (a *Autoplius) ProcessQuery(q *Query, b *Browser) ([]Ad, error) {
	// do something
	log.Println("[Autoplius] Processing query:", q.Name)
	b.ResetTimer()
	doc, err := b.LoadPage(q.FirstURL)
	// b.CancelTimer()
	if err != nil {
		return nil, err
	}
	pageID := 1
	ads, _ := a.parseAnnouncements(doc)
	log.Printf("[Autoplius] Loaded page #%v ads:%v", pageID, len(ads))

	// check next page
	nextURL := a.parseNextPageURL(doc)
	for nextURL != "" {
		pageID++
		//
		b.ResetTimer()
		doc, err := b.LoadPage("https://autoplius.lt" + nextURL)
		// b.CancelTimer()
		if err != nil {
			log.Println("[Autoplius] Error loading page:", err)
			return nil, err
		}
		// store
		pageAds, _ := a.parseAnnouncements(doc)
		ads = append(ads, pageAds...)
		// log
		log.Printf("[Autoplius] Loaded page #%v ads:%v", pageID, len(pageAds))
		// for _, an := range pageAnnouncements {
		// 	anID++
		// 	log.Printf("#%d %v '%v' '%v' %v", anID, an.Title, an.TitleParams, an.Params, an.Price)
		// }
		nextURL = a.parseNextPageURL(doc)
	}
	return ads, nil
}

// getNextURL returns the next page URL
func (a *Autoplius) parseNextPageURL(doc *goquery.Document) string {
	nextURL, _ := doc.Find(".page-navigation-container").Find("a.next").Attr("href")
	return nextURL
}

// ParseAnnouncementsFromReader parses the HTML document from the given reader
func (a *Autoplius) parseAnnouncements(doc *goquery.Document) ([]Ad, error) {
	// // Load the HTML document
	// doc, err := goquery.NewDocumentFromReader(r)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	announcements := []Ad{}

	// Find the review items
	doc.Find(".announcement-item").Each(func(i int, s *goquery.Selection) {
		// get the ad URL
		// ad_url := s.AttrOr("href", "")
		ad_url, _ := s.Attr("href")
		// log.Println("ad_url: ", ad_url)
		// get the title
		title := s.Find(".announcement-title").Text()
		title = strings.TrimSpace(title)
		// get title params
		titleParams := s.Find(".announcement-title-parameters").Text()
		titleParams = strings.TrimSpace(titleParams)
		titleParams = removeMultipleSpaces(titleParams)
		// get announcement photo url
		// imageURL, _ := s.Find(".announcement-media").Html()
		// imageURL, _ := s.Find(".announcement-photo").Find("img").Html()
		src, _ := s.Find(".announcement-media img").Attr("src")
		dataSrc, _ := s.Find(".announcement-media img").Attr("data-src")
		imageURL := dataSrc
		if dataSrc == "" {
			imageURL = src
		}
		// log.Println("imageURL: ", imageURL)

		// get the price
		price := s.Find(".announcement-pricing-info").Text()
		price = strings.TrimSpace(price)
		price = removeMultipleSpaces(price)
		price = removeSpacesBetweenDigits(price)
		// get parameters block
		params := s.Find(".announcement-parameters-block").Text()
		params = strings.TrimSpace(params)
		params = removeMultipleSpaces(params)

		ad := Ad{
			Title:      title,
			Subtitle:   titleParams,
			Price:      price,
			Other:      params,
			ImageURLs:  []string{imageURL},
			ScrapeTime: time.Now(),
			AdURL:      ad_url,
		}
		ad.AdvertiserAdID = ad.ImageURLs[0]
		announcements = append(announcements, ad)
	})

	return announcements, nil
}

// GetName returns the advertiser name
func (a *Autoplius) GetName() string {
	return a.Name
}

// GetFormURL returns the advertiser form URL
func (a *Autoplius) GetFormURL() string {
	return a.FormURL
}
