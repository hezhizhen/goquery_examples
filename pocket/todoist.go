package main

import (
	"github.com/PuerkitoBio/goquery"
)

func handleTodoist(s *goquery.Selection) (string, string) {
	post := s.Find("h2.tdb-article-slat__title a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	title, exist := post.Attr("title")
	if !exist {
		panic("missing title")
	}
	return title, url
}
