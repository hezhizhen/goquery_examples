package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleReadErn(url string, p Pocket) {
	var urls []string
	for {
		fmt.Println(url)
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("article.post")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.post-title.entry-title a").Attr("href")
			if !exist {
				fmt.Println(s.Html())
				panic("missing url")
			}
			urls = append(urls, href)
		})
		nexts := doc.Find("div.pagenav a")
		next := ""
		exist := false
		nexts.Each(func(i int, s *goquery.Selection) {
			if !s.HasClass("number") && s.Text() == ">" {
				next, exist = s.Attr("href")
				return
			}
		})
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
