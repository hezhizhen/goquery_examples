package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleLepture(url string, p Pocket) {
	var urls []string
	suffix := "/archive/"
	for {
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.item.Article")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		prev, exist := doc.Find("div.navigation.color a.prev").Attr("href")
		if !exist {
			break
		}
		suffix = prev
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
