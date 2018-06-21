package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleAirOnG(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("article.post")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		older, exist := doc.Find("nav.pagination a.older-posts").Attr("href")
		if !exist {
			break
		}
		suffix = older
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
