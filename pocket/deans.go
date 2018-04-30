package main

import "github.com/PuerkitoBio/goquery"

func handleDeansSite(s *goquery.Selection) (string, string) {
	post := s.Find("div.post_title a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
