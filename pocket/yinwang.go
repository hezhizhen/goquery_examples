package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleYinWang(p Pocket, skip bool) {
	if skip {
		fmt.Println("Skip: http://www.yinwang.org")
	}
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
		p.Add(postURL)
		fmt.Printf("Successfully saved article to pocket whose title is: %s\n", s.Find("a").Text())
	})
	fmt.Println("Saved all posts from blog http://www.yinwang.org/")
}
