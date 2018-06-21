package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleLiYang85(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.fusion-post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("a.pagination-next").Attr("href")
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
