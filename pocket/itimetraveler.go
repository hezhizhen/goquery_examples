package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleITimeTraveler(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url + "/archives/")
	handleError(err)
	list := doc.Find("article.archive-article.archive-type-post")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+strings.TrimPrefix(href, ".."))
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
