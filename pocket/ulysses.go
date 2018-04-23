package main

import "github.com/PuerkitoBio/goquery"

func handleUlysses(s *goquery.Selection) (string, string) {
	post := s.Find("header.entry-header h1.entry-title a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
