package main

import "github.com/PuerkitoBio/goquery"

func handleCatCoder(s *goquery.Selection) (string, string) {
	title := s.Find("h1.post-title a")
	url, exist := title.Attr("href")
	if !exist {
		panic("missing url")
	}
	return title.Text(), url
}
