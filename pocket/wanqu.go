package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleWanQu(url string, p Pocket) {
	var urls []string
	var suffix string
	count := 1
	for {
		fmt.Println(url + suffix)
		doc, err := goquery.NewDocument(url + suffix)
		handleError(err)
		list := doc.Find("li.list-group-item h2.list-title")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, url+trimSuffix(href))
		})
		prev, exist := doc.Find("nav li.previous a").Attr("href")
		if !exist {
			break
		}
		if count == 100 {
			p.AddMultiple(urls)
			for i := range urls {
				fmt.Printf("%3d: %s\n", i+1, urls[i])
			}
			urls = []string{}
			count = 1
		} else {
			count++
		}
		suffix = trimSuffix(prev)
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}

func trimSuffix(href string) string {
	parts := strings.Split(href, "?s=")
	return parts[0]
}
