package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleGravitySworm(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div#content article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("div.meta div.share div.share-links a.share-google").Attr("href")
			if !exist {
				panic("missing url")
			}
			parts := strings.Split(href, "url=")
			urls = append(urls, parts[1])
		})
		next, exist := doc.Find("div#pagination a.next").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
