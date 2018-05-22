package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleAppShere(url string, p Pocket) {
	doc, err := goquery.NewDocument(url + "/archive")
	handleError(err)
	var urls []string
	list := doc.Find("div.archives li")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Println(i+1, ":", urls[i])
	}
}
