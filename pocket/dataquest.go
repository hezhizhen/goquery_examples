package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleDataQuest(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("article.post-card")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a.post-card-content-link").Attr("href")
		if !exist {
			panic("missing url")
		}
		href = strings.TrimPrefix(href, "/blog")
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
