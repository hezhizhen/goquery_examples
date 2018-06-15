package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleXiaoRui(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div.article.well.clearfix")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("div.title-article a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("a#load-more").Attr("href")
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
