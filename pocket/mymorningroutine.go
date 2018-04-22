package main

import (
	"github.com/PuerkitoBio/goquery"
)

func handleMyMorningRoutine(s *goquery.Selection) (string, string) {
	block := s.Find("a.u-block")
	title, exist := block.Attr("title")
	if !exist {
		panic("missing title")
	}
	post, exist := block.Attr("href")
	if !exist {
		panic("missing url")
	}
	return title, post
}
