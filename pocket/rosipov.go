package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleROsipov(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url + "/blog/archive/")
	handleError(err)
	list := doc.Find("div#blog-archives li")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
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
