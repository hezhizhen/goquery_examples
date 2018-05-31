package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleEverythingCLI(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div#main article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("div.entry-content h2.blog-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		pages := doc.Find("ul.pag li")
		current := -1
		noMore := true
		pages.Each(func(i int, s *goquery.Selection) {
			if s.HasClass("current") {
				current = i
			}
			if current != -1 && i != current {
				next, exist := s.Find("a").Attr("href")
				if !exist {
					panic("missing url for next")
				}
				url = next
				noMore = false
				return
			}
		})
		if noMore {
			break
		}
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
