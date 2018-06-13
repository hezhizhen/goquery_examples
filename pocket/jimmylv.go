package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJimmyLv(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url + "/pages/archive/")
	handleError(err)
	list := doc.Find("div#page-content li")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("a.pjaxlink").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
