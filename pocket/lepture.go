package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleLepture() {
	url := "https://lepture.com"
	part := "/archive/"
	for {
		doc, err := goquery.NewDocument(url + part)
		handleError(err)
		list := doc.Find("div[id][class]")
		list.Each(func(i int, s *goquery.Selection) {
			post, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			saveToPocket(url + post)
			fmt.Printf("Successfully saved article to pocket whose title is: %s\n", s.Find("h3").Text())
		})
		prev, exist := doc.Find("div.navigation.color").Find("a.prev").Attr("href")
		if !exist {
			break
		}
		part = prev
	}
}
