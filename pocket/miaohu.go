package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleMiaoHu() {
	url := "https://miao.hu/"
	doc, err := goquery.NewDocument(url)
	handleError(err)
	list := doc.Find("li.mv2")
	list.Each(func(i int, s *goquery.Selection) {
		post, exist := s.Find("a").Attr("href")
		if !exist {
			panic("missing url")
		}
		saveToPocket(post)
		time := s.Find("time").Text()
		title := strings.TrimSpace(s.Text())
		title = strings.TrimPrefix(title, time)
		fmt.Printf("Successfully saved article to pocket whose title is: %s\n", strings.TrimSpace(title))
	})
}
