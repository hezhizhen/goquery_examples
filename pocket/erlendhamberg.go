package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleErlendHamberg(url string, p Pocket) {
	var urls []string
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("div#posts div.postitem")
	list.Each(func(i int, s *goquery.Selection) {
		post := s.Find("div.postsummary div.posttitle a")
		href, exist := post.Attr("href")
		if !exist {
			panic("missing url")
		}
		href = strings.TrimPrefix(href, ".")
		urls = append(urls, url+href)
	})
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Println(i+1, ":", urls[i])
	}
}
