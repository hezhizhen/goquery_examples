package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleYinWangLofter() {
	url := "http://yinwang0.lofter.com/?page=%d"
	page := 1
	for {
		doc, err := goquery.NewDocument(fmt.Sprintf(url, page))
		handleError(err)
		list := doc.Find("div.m-post.m-post-txt")
		list.Each(func(i int, s *goquery.Selection) {
			title := s.Find("h2.ttl")
			postURL, exist := title.Find("a").Attr("href")
			if !exist || postURL == "" {
				panic("missing url for post")
			}
			saveToPocket(postURL)
			fmt.Printf("Successfully saved article to pocket whose title is: %s\n", title.Text())
		})
		if list.Length() < 10 {
			break
		}
		page++
	}
}
