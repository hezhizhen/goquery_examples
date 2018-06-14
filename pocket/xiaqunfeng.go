package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleXiaQunFeng(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("article.article.article-type-post.article-index")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("header.article-header h1 a.article-title").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("nav#page-nav a.extend.next").Attr("href")
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
