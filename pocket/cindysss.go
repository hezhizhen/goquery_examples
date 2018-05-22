package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleCindysss(url string, p Pocket) {
	var urls []string
	for {
		doc, err := goquery.NewDocument(url)
		handleError(err)
		list := doc.Find("div#content article[id]")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h1.entry-title a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("nav#nav-below div.nav-previous a").Attr("href")
		if !exist {
			break
		}
		url = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Println(i+1, ":", urls[i])
	}
}
