package main

import "github.com/PuerkitoBio/goquery"

func handleJez(s *goquery.Selection) (string, string) {
	post := s.Find("h3.entry-title a")
	title, exist := post.Attr("title")
	if !exist {
		panic("missing title")
	}
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return title, url
}
