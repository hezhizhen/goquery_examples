package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleSheepBao(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.posts-list article.post-preview")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("ul.pager.main-pager li.next a").Attr("href")
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
