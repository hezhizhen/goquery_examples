package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleJosuiWritings(p Pocket, url string) {
	nextPart := "/"
	page := 1
	for {
		doc, err := goquery.NewDocument(url + nextPart)
		handleError(err)

		list := doc.Find("div.blog div.content article")
		urls := []string{}
		titles := []string{}
		list.Each(func(i int, s *goquery.Selection) {
			title := s.Find("h2.article-title a")
			post, exist := title.Attr("href")
			if !exist {
				panic("url is missing")
			}
			titles = append(titles, title.Text())
			urls = append(urls, url+post)
		})
		p.AddMultiple(urls)
		fmt.Printf("Haved saved %d articles in page %d:\n", len(titles), page)
		fmt.Println(strings.Join(titles, "\n"))
		next := doc.Find("nav.pagination a.pagination-next")
		nextURL, exist := next.Attr("href")
		if !exist {
			break
		}
		nextPart = nextURL
		page++
	}
}
