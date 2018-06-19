package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleGo101(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("ul li")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, strings.TrimSuffix(url, "101.html")+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
