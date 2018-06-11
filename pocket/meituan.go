package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMeiTuan(url string, p Pocket) {
	var suffix string
	var doc *goquery.Document
	var err error
	var urls []string
	for {
		fmt.Println(url + suffix)
		doc, err = goquery.NewDocument(url + suffix)
		handleError(err)
		more, exist := doc.Find("footer.more a").Attr("href")
		if !exist {
			break
		}
		suffix = more
	}
	list := doc.Find("article.post")
	list.Each(func(i int, s *goquery.Selection) {
		href, exist := s.Find("header.post-title a").Attr("href")
		if !exist {
			panic("missing url")
		}
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
