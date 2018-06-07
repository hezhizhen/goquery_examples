package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handlePopoZhu(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.archive_title")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("div.archive_page a.extend.next").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
