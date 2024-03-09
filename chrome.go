package carbot

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type Browser struct {
	debugFilename string
	debugURL      string
	// browser context
	ctx           context.Context
	cancelTimer   context.CancelFunc // cancel timeout
	cancelContext context.CancelFunc // cancel timeout
	cancelRemote  context.CancelFunc // cancel timeout
}

// NewBrowser creates a new Browser instance
func NewBrowser() *Browser {
	b := &Browser{
		debugFilename: ".tmp/chromium_debug.txt",
		debugURL:      "",
	}
	b.debugURL = b.getDebugURL(b.debugFilename)
	b.ConnectBroser()
	return b
}

// LoadPage loads the page
func (b *Browser) LoadPage(url string) (*goquery.Document, error) {
	var htmlContent string
	err := chromedp.Run(b.ctx,
		chromedp.Navigate(url),
		// chromedp.WaitVisible(`body`),
		// chromedp.WaitVisible(`footer`),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)
	if err != nil {
		return nil, err
	}
	return GoQueryFromString(htmlContent)
}

// get Context returns the browser context
func (b *Browser) ConnectBroser() {
	// allocates a new remote browser context to use with already running browser
	var allocCtx context.Context
	allocCtx, b.cancelRemote = chromedp.NewRemoteAllocator(context.Background(), b.debugURL)
	// defer cancel()

	// create a new context
	b.ctx, b.cancelContext = chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
}

// CloseBrowser closes the browser
func (b *Browser) CloseBrowser() {
	b.cancelContext() // close tab
	b.cancelRemote()
	b.cancelTimer()
}

// ResetTimer resets the timer
func (b *Browser) ResetTimer() {
	b.ctx, b.cancelTimer = context.WithTimeout(b.ctx, 5*time.Second)
	// defer cancel()
}

// // CancelTimer cancels the timer
// func (b *Browser) CancelTimer() {
// 	b.cancelTimer()
// }

func (b *Browser) getDebugURL(url string) string {
	// Open the file
	file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new buffered reader
	reader := bufio.NewReader(file)

	// Read only the first line of the file
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trim the newline character from the end of the line
	line = line[:len(line)-1]

	return line
}

// getNextURL returns the next page URL
// // this goes to autoplius.go
// func ParseNextPageURL(doc *goquery.Document) string {
// 	nextURL, _ := doc.Find(".page-navigation-container").Find("a.next").Attr("href")
// 	return nextURL
// }

// JQueryDocumentFromString returns a goquery document from the given string
func GoQueryFromString(s string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(s))
}
