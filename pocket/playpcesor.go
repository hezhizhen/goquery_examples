package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handlePlayPCEsor(url string, p Pocket) {
	for {
		var urls []string
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("article.post-outer-container")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h3.post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		more, exist := doc.Find("div.blog-pager a.blog-pager-older-link").Attr("href")
		if !exist {
			break
		}
		url = more
		p.AddMultiple(urls)
		for i := range urls {
			fmt.Printf("%2d: %s\n", i+1, urls[i])
		}
	}
	fmt.Println("done")
}
