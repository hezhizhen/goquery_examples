package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleAgilebits(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div.template-blog article.post-entry")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.post-title.entry-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		nav := doc.Find("nav.pagination a")
		var next string
		nav.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "›" {
				href, exist := s.Attr("href")
				if !exist {
					panic("missing url for next")
				}
				next = href
				return
			}
		})
		if next == "" {
			break
		}
		url = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
