package main

// https://www.zenrows.com/blog/web-scraping-golang#scraping-dynamic-content

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	// "net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	// "github.com/gocolly/colly/v2"
	// "https://github.com/juju/persistent-cookiejar"
)

func main() {
	chromeDPScrape()
	// netHttpScrape()
	// collyScrape()
}

func chromeDPScrape() {
	// chromium-browser --remote-debugging-port=9222
	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/devtools/browser/2ed26278-89fd-4105-b8b5-9ce7f15f7b80")

	// Create a new allocator context for custom options
	// chromiumPath := "/usr/bin/chromium-browser"
	// allocCtx, cancel := chromedp.NewExecAllocator(context.Background(),
	// 	append(chromedp.DefaultExecAllocatorOptions[:],
	// 		// Set the path to the Chromium binary
	// 		chromedp.ExecPath(chromiumPath),
	// 	)...,
	// )
	defer cancel()

	// Create a new context
	// ctx, cancel := chromedp.NewContext(context.Background())
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	// defer cancel()

	// Set a timeout to prevent hanging
	ctx, cancel = context.WithTimeout(ctx, 35*time.Second)
	defer cancel()

	var htmlContent string
	var b1 []byte
	err := chromedp.Run(ctx,
		// chromedp.Emulate(device.Reset),
		// chromedp.Emulate(device.GalaxyTabS4landscape),
		// chromedp.Emulate(device.JioPhone2),
		// network.Enable(), // Enable network domain for cookie manipulation
		chromedp.Navigate(`https://autoplius.lt/skelbimai/naudoti-automobiliai?make_id=104`),
		// Wait for the footer element to ensure the page has loaded. You might need to adjust this.
		// chromedp.WaitVisible(`body`),
		// chromedp.WaitVisible(`footer`),
		// Extract the HTML of the rendered page
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
		chromedp.CaptureScreenshot(&b1),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("screenshot1.png", b1, 0o644); err != nil {
		log.Fatal(err)
	}

	fmt.Println(htmlContent)
	os.WriteFile("body.html", []byte(htmlContent), 0644)
	// Now you can use htmlContent with goquery or Colly's NewDocumentFromReader to parse and scrape the needed data.

	// // COOKIES
	// // Example of retrieving all cookies
	// var cookies []*network.Cookie
	// err = chromedp.Run(ctx,
	// 	chromedp.ActionFunc(func(ctx context.Context) error {
	// 		var err error
	// 		cookies, err = network.GetAllCookies().Do(ctx)
	// 		return err
	// 	}),
	// )

	// if err != nil {
	// 	log.Fatalf("Failed to retrieve cookies: %v", err)
	// }

	// // Print the cookies
	// for _, cookie := range cookies {
	// 	fmt.Printf("Cookie: %s=%s; Domain: %s; Expires: %s\n", cookie.Name, cookie.Value, cookie.Domain, time.Unix(cookie.Expires, 0).Format(time.RFC1123))
	// }
}

func netHttpScrape() {
	// The URL of the page to scrape
	url := "https://autoplius.lt/skelbimai/naudoti-automobiliai?make_id=104"

	// Fetch the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Parse the page body with goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find and print car titles (adjust the selector as needed)
	doc.Find(".car-title").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Car %d: %s\n", i+1, title)
	})
}

func collyScrape() {
	c := colly.NewCollector(
		//colly.CacheDir("./_instagram_cache/"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0"),
		// colly.Debugger(&debug.LogDebugger{}),
	)

	j, err := cookiejar.New(&cookiejar.Options{})
	// j, err := cookiejar.New(&cookiejar.Options{Filename: "cookie.db"})
	if err == nil {
		c.SetCookieJar(j)
	}

	// // find and visit all links: class="next" rel="next"
	// c.OnHTML("a.next", func(e *colly.HTMLElement) {
	// 	// c.OnHTML("a.next[rel=next]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href") // Get the link in the href attribute
	// 	if link != "" {
	// 		e.Request.Visit(link)
	// 	} else {
	// 		log.Println("No more pages to visit")
	// 	}
	// })

	// c.OnHTML("html", func(e *colly.HTMLElement) {
	// 	log.Println("OnHTML")
	// 	d := c.Clone()
	// 	d.OnResponse(func(r *colly.Response) {
	// 		body := string(r.Body)
	// 		log.Println("Visited", r.Request.URL, len(body))
	// 		// requestIds = queryIdPattern.FindAll(r.Body, -1)
	// 		// requestID = string(requestIds[1][9:41])
	// 	})
	// })

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting:", r.URL)
		r.Headers.Set("Referer", "https://autoplius.lt")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Host", "autoplius.lt")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
		r.Headers.Set("Sec-Fetch-Site", "none")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0")
		r.Headers.Set("Cookie", "name=value")
	})

	c.OnResponse(func(r *colly.Response) {
		body := string(r.Body)
		// log.Println("Visited", r.Request.URL, len(body))
		log.Println("Visited", r.Request.URL, body)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode)
		log.Println("Request Headers:", r.Request.Headers)
		log.Println("Body:", string(r.Body))
		log.Println("Error:", err)
		// save r.Body to file
		os.WriteFile("body.html", r.Body, 0644)

	})

	// // 15
	// c.Visit("https://autoplius.lt/skelbimai/naudoti-automobiliai?make_date_from=2016&make_date_to=&sell_price_from=&sell_price_to=30000&engine_capacity_from=&engine_capacity_to=&power_from=&power_to=&kilometrage_from=&kilometrage_to=&qt=&has_damaged_id=10924&steering_wheel_id=10922&category_id=2&make_id_list=72&make_id%5B72%5D=19859&slist=2207751801")
	// 69
	err = c.Visit("https://m.autoplius.lt/skelbimai/naudoti-automobiliai?make_id=104")
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return
	}
	// c.Visit("https://autoplius.lt/skelbimai/naudoti-automobiliai?make_date_from=2010&make_date_to=&sell_price_from=&sell_price_to=30000&engine_capacity_from=&engine_capacity_to=&power_from=&power_to=&kilometrage_from=&kilometrage_to=&qt=&has_damaged_id=10924&steering_wheel_id=10922&category_id=2&make_id_list=72&make_id%5B72%5D=19859&slist=2207751801")
}
