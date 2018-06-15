package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMyCodeSmells(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.post")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("div.post-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next := ""
		exist := false
		pages := doc.Find("a.nav-btn.newer-btn")
		pages.Each(func(i int, s *goquery.Selection) {
			if s.Find("i").HasClass("fa-long-arrow-right") {
				next, exist = s.Attr("href")
			}
		})
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
