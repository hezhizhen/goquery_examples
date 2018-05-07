package main

import "github.com/PuerkitoBio/goquery"

func handleYuelong(s *goquery.Selection) (string, string) {
	post := s.Find("h2 a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
