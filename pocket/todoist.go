package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleTodoist(p Pocket, url string) {
	for {
		urls := []string{}
		doc, err := goquery.NewDocument(url)
		handleError(err)
		doc.Find("article.tdb-article-slat").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h2.tdb-article-slat__title a")
			post, exist := title.Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, post)
		})
		p.AddMultiple(urls)
		fmt.Printf("Successfully saved %d articles to pocket in %s\n", len(urls), url)
		older, exist := doc.Find("div.tdb-pagination-holder a.next.page-numbers.nav__action").Attr("href")
		if !exist {
			break
		}
		url = older
	}
}
