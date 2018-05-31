package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJust4Fun(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div#content div.article")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h1 a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("div.pagination li.next a").Attr("href")
		if !exist {
			break
		}
		if next == "#" {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
