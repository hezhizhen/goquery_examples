package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJamesStuber(url string, p Pocket) {
	ss := []string{
		"/articles/",
		"/booknotes/",
	}
	for _, s := range ss {
		var urls []string
		doc, err := goquery.NewDocument(url + s)
		handleError(err)
		list := doc.Find("div.posts ul li.post-item")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		p.AddMultiple(urls)
		fmt.Println("Posts from", url+s)
		for i := range urls {
			fmt.Println(i+1, ":", urls[i])
		}
	}
}
