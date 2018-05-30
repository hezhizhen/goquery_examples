package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMacStories(url string, p Pocket) {
	var urls []string
	total := 0
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div.posts article.post")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("header.post-header h1.post-title a.post-link").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("nav.page-navigation a.next").Attr("href")
		if !exist {
			break
		}
		url = next
		total++
		if total > 123 {
			break
		}
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%4d: %s\n", i+1, urls[i])
	}
}
