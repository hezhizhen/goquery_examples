package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleYinWangLofter(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("div.m-postlst div.m-post.m-post-txt")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h2.ttl a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		next, exist := doc.Find("div.m-pager.m-pager-idx.box a.next").Attr("href")
		if !exist {
			break
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
