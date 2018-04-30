package main

import "github.com/PuerkitoBio/goquery"

func handleUneeWang(s *goquery.Selection) (string, string) {
	post := s.Find("div.post-main h3.post-title a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
