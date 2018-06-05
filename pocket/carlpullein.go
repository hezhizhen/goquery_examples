package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleCarlPullein(url string, p Pocket) {
	var urls []string
	suffix := "/blog"
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div#content div.main-content article.hentry")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("div.post h1.entry-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("nav.page.pagination li.next a#nextLink").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
