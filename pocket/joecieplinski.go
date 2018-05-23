package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJoecieplinski(url string, p Pocket) {
	var urls []string
	suffix := "/blog"
	for {
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("main.content article.preview")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h1.post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("nav.pagination a.older-posts").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Println(i+1, ":", urls[i])
	}
}
