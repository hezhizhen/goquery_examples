package main

import "github.com/PuerkitoBio/goquery"

func handleAsianEfficiency(s *goquery.Selection) (string, string) {
	post := s.Find("h1 a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
