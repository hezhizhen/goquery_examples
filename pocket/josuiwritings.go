package main

import (
	"github.com/PuerkitoBio/goquery"
)

func handleJosuiWritings(s *goquery.Selection) (string, string) {
	title := s.Find("h2.article-title a")
	post, exist := title.Attr("href")
	if !exist {
		panic("url is missing")
	}
	return title.Text(), post
}
