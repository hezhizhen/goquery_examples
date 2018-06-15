package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleCalhoun(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.block.pt-6.pb-0")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a.no-underline").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("div.inline-flex a.rounded-r").Attr("href")
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
