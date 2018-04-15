package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleLiQi() {
	p := NewPocket()
	url := "http://liqi.io/"
	total := 0
	urls := []string{}
	for {
		doc, err := goquery.NewDocument(url)
		handleError(err)
		doc.Find("article[id]").Each(func(i int, s *goquery.Selection) {
			post, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, post)
			total++
		})
		p.AddMultiple(urls)
		prev, exist := doc.Find("div.nav-previous a").Attr("href")
		if !exist {
			break
		}
		fmt.Printf("Successfully saved %d articles to pocket in total\n", total)
		url = prev
	}
}
