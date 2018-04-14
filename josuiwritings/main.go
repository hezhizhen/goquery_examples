package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// read the article about how to get access token: http://www.cnblogs.com/febwave/p/4242333.html
	// get access token
	type auth struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
	}
	f, err := os.Open("auth.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	a := auth{}
	if err := json.Unmarshal(bs, &a); err != nil {
		panic(err)
	}

	// find all articles
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
