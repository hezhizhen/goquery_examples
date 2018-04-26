package main

import "github.com/PuerkitoBio/goquery"

func handleStartUpNextDoor(s *goquery.Selection) (string, string) {
	title := s.Find("h2.post-title a")
	url, exist := title.Attr("href")
	if !exist {
		panic("missing url")
	}
	return title.Text(), url
}
