package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleLeetcodeArticle(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("a.list-group-item")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, "https://leetcode.com"+href)
		})
		next, exist := doc.Find("nav li.next a").Attr("href")
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
