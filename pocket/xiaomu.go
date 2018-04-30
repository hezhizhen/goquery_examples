package main

import "github.com/PuerkitoBio/goquery"

func handleXiaomu(s *goquery.Selection) (string, string) {
	post := s.Find("h2.post-title a")
	url, exist := post.Attr("href")
	if !exist {
		panic("missing url")
	}
	return post.Text(), url
}
