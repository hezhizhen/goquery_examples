package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleChewxy(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("article.post-preview")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		older, exist := doc.Find("li.next a").Attr("href")
		if !exist {
			break
		}
		suffix = older
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
