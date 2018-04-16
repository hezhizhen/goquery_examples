package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleJannerChang(p Pocket) {
	url := "http://jannerchang.bitcron.com"
	part := "/"
	for {
		urls := []string{}
		doc, err := goquery.NewDocument(url + part)
		handleError(err)
		doc.Find("div.post_in_list").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h2 a")
			post, exist := title.Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+post)
		})
		p.AddMultiple(urls)
		fmt.Printf("Successfully saved %d articles to pocket in %s\n", len(urls), url+part)
		next, exist := doc.Find("div.paginator.pager.pagination a.btn.next.older-posts.older_posts").Attr("href")
		if !exist {
			break
		}
		part = next
	}
}
