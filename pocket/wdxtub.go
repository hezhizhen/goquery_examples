package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleWdxtub(url string, p Pocket) {
	var urls []string
	suffix := "/archives/"
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("article.post.post-type-normal")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.post-title a.post-title-link").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("nav.pagination a.extend.next").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
