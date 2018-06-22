package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleAlgorithmNote(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("div.l")
	list.Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, t *goquery.Selection) {
			href, exist := t.Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, strings.TrimSuffix(url, "index.html")+href)
		})
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
