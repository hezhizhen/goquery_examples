package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleZFanW2(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div._x2")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a._x3").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		found := false
		next := ""
		pages := doc.Find("div._xg").Children()
		pages.Each(func(i int, s *goquery.Selection) {
			if s.HasClass("_xh") {
				found = true
				return
			}
			if found {
				href, exist := s.Attr("href")
				if !exist {
					panic("missing url for next page")
				}
				next = href
				found = false
			}
		})
		if next == "" {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
