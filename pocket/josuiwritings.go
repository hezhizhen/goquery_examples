package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJosuiWritings() {
	p := NewPocket()
	blogURL := "http://blog.josui.me"
	nextURL := "/archives/"
	exist := true
	for {
		doc, err := goquery.NewDocument(blogURL + nextURL)
		handleError(err)

		list := doc.Find("div.archive").Find("div.post.archive")
		list.Each(func(i int, s *goquery.Selection) {
			title := s.Find(".archive-title a")
			url, exist := title.Attr("href")
			if !exist || url == "" {
				panic("url is missing")
			}
			url = fmt.Sprintf("http://blog.josui.me%s", url)
			p.Add(url)
			fmt.Printf("Successfully saved article to pocket whose title is: %s\n", title.Text())
		})
		next := doc.Find("a.pagination-next")
		nextURL, exist = next.Attr("href")
		if !exist {
			break
		}
	}
}
