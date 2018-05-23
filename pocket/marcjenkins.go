package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleMarcJenkins(url string, p Pocket) {
	var urls []string
	current := url + "/blog"
	for {
		doc, err := goquery.NewDocument(current)
		handleError(err)
		list := doc.Find("ol.latest-posts li")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("article.article h2 a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, href)
		})
		nav := doc.Find("nav.blog-nav a")
		var next string
		exist := false
		nav.Each(func(i int, s *goquery.Selection) {
			if s.Text() == "Next" {
				next, exist = s.Attr("href")
				if !exist {
					panic("missing url for next page")
				}
			}
		})
		if !exist {
			break
		}
		current = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
