package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMikeWootc(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("div#wrapper p")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if exist {
			urls = append(urls, url+href)
		}
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
