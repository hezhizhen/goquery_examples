package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleYinWang() {
	url := "http://www.yinwang.org"
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("li.list-group-item.title")
	list.Each(func(i int, s *goquery.Selection) {
		postURL, exist := s.Find("a").Attr("href")
		if !exist || url == "" {
			panic("missing url for post")
		}
		postURL = fmt.Sprintf("%s%s", url, postURL)
		saveToPocket(postURL)
		fmt.Printf("Successfully saved article to pocket whose title is: %s\n", s.Find("a").Text())
	})
}
