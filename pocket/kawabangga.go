package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleKawabangga(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url + "/all-posts")
	handleError(err)
	list := doc.Find("ul.car-list li")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
