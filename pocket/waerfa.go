package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleWaerfa(s *goquery.Selection) (string, string) {
	title := s.Find("h2.Article__title a")
	url, exist := title.Attr("href")
	if !exist {
		panic("missing url")
	}
	text := strings.TrimSpace(title.Text())
	return text, url
}
