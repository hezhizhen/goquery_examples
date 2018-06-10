package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMissPower(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div.article div.note-container")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Attr("data-url")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("div.paginator span.next a").Attr("href")
		if !exist {
			break
		}
		url = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
