package main

import (
	"github.com/PuerkitoBio/goquery"
)

func handleYinWang(s *goquery.Selection) (string, string) {
	post := s.Find("a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
