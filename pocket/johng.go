package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJohng(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.entry-title a[rel=bookmark]").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("nav.navigation.pagination div.nav-links a.next.page-numbers").Attr("href")
		if !exist {
			break
		}
		url = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
