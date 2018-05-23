package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleNladuo(url string, p Pocket) {
	var urls []string
	var suffix string
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("ul.home.post-list li.post-list-item")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a.post-title-link").Attr("href")
			if !exist {
				fmt.Println(s.Html())
				panic("missing url")
			}
			if strings.HasPrefix(href, "../../") {
				href = strings.TrimPrefix(href, "../../")
			}
			urls = append(urls, url+href)
		})
		next, exist := doc.Find("div.paginator a.next").Attr("href")
		if !exist {
			break
		}
		if strings.HasPrefix(next, "../") {
			next = strings.Replace(next, "../", "page/", -1)
		}
		suffix = next
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%2d: %s\n", i+1, urls[i])
	}
}
