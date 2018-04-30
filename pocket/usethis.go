package main

import "github.com/PuerkitoBio/goquery"

func handleUseThis(s *goquery.Selection) (string, string) {
	post := s.Find("h3 a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
