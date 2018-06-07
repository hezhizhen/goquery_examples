package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleBitsplitting(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div#content article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h1.entry-title a[rel=bookmark]").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		older, exist := doc.Find("nav.navigation div.nav-previous a").Attr("href")
		if !exist {
			break
		}
		url = older
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
