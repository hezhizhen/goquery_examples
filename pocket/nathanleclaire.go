package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleNathanLeclaire(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("tr.post-row")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
