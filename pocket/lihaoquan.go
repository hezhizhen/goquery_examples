package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleLiHaoQuan(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url + "archive")
	handleError(err)
	list := doc.Find("h5")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+strings.TrimPrefix(href, "/"))
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
