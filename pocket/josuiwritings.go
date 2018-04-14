package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func handleJosuiWritings(a Auth) {
	blogURL := "http://blog.josui.me"
	nextURL := "/archives/"
	exist := true
	for {
		doc, err := goquery.NewDocument(blogURL + nextURL)
		if err != nil {
			panic(err)
		}
		list := doc.Find("div.archive").Find("div.post.archive")
		list.Each(func(i int, s *goquery.Selection) {
			title := s.Find(".archive-title a")
			url, exist := title.Attr("href")
			if !exist || url == "" {
				panic("url is missing")
			}
			url = fmt.Sprintf("http://blog.josui.me%s", url)
			body := []byte(fmt.Sprintf(`{
				"url": "%s",
				"consumer_key": "%s",
				"access_token": "%s"
			}`, url, a.ConsumerKey, a.AccessToken))
			req, err := http.Post("https://getpocket.com/v3/add", "application/json", bytes.NewReader(body))
			if err != nil {
				panic(err)
			}
			if req.StatusCode != 200 {
				panic("fail to save the article to pocket")
			}
			fmt.Printf("Successfully saved article to pocket whose title is: %s\n", title.Text())
		})
		next := doc.Find("a.pagination-next")
		nextURL, exist = next.Attr("href")
		if !exist {
			break
		}
	}
}
