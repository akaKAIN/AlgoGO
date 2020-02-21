package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goRoutine/MyLogger"
	"log"
	"net/http"
	"os"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://infodev.ru")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("a.nav-link").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("p").Text()
		link, ok := s.Attr("href")
		if !ok {
			link = "нет ссылки"
		}
		fmt.Printf("меню: %s - %s\n", title, link)
	})
}

func main() {
	ExampleScrape()
	_, err := os.OpenFile("awdawdaw", os.O_APPEND, 072)
	MyLogger.ForError(err)
	MyLogger.ForError(nil)
}
