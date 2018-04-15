package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleLeetcodeArticle(p Pocket) {
	url := "https://leetcode.com/articles/?page=%d"
	page := 1
	for {
		doc, err := goquery.NewDocument(fmt.Sprintf(url, page))
		handleError(err)
		list := doc.Find("a.list-group-item")
		list.Each(func(i int, s *goquery.Selection) {
			postURL, exist := s.Attr("href")
			if !exist {
				panic("missing url")
			}
			p.Add(fmt.Sprintf("https://leetcode.com%s", postURL))
			fmt.Printf("Successfully saved article to pocket whose title is: %s\n", strings.TrimSpace(s.Find("h4.media-heading").Text()))
		})
		if list.Length() < 10 {
			break
		}
		page++
	}
}
