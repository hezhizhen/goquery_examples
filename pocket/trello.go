package main

import (
	"github.com/PuerkitoBio/goquery"
)

func handleTrello(s *goquery.Selection) (string, string) {
	title := s.Find("div h2 a")
	url, exist := title.Attr("href")
	if !exist {
		panic("missing url")
	}
	return title.Text(), url
}
