package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleStevenOcchipinti(url string, p Pocket) {
	var urls []string
	suffix := "/archives/"
	doc, err := goquery.NewDocument(url + suffix)
	handleError(err)
	list := doc.Find("article")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("h1 a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
