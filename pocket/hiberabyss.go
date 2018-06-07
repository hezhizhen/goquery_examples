package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleHiberAbyss(url string, p Pocket) {
	var suffix string
	var urls []string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.content_container div.post")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h1.post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("nav.page-navigator a.extend.next").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
