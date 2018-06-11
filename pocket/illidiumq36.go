package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleIllidiumq36(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div#content h2.entry-title")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a[rel=bookmark]").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		older, exist := doc.Find("div.navigation div.nav-previous a").Attr("href")
		if !exist {
			break
		}
		url = older
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
