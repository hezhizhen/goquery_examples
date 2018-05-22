package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func handleIBTSAT(url string, p Pocket) {
	suffices := []string{
		"/categories/toefl-listening.html",
	}
	for _, s := range suffices {
		var urls, titles []string
		currentURL := url + s
		for {
			fmt.Println("new doc for " + currentURL)
			doc, err := goquery.NewDocument(currentURL)
			handleError(err)
			list := doc.Find("div.col-sm-9 ul.list-unstyled.list-main li")
			list.Each(func(i int, s *goquery.Selection) {
				href, exist := s.Find("a").Attr("href")
				if !exist {
					panic("missing url")
				}
				title := s.Find("a h4").Text()
				titles = append(titles, title)
				urls = append(urls, url+href)
			})
			end := false
			doc.Find("nav ul.pagination.justify-content-end li.page-item a.page-link").Each(func(i int, s *goquery.Selection) {
				end = true
				if s.Text() == "下一页" {
					u, exist := s.Attr("href")
					if !exist {
						panic("missing url for next page")
					}
					if u == "javascript:void(0)" {
						return
					}
					currentURL = url + u
					end = false
				}
			})
			if end {
				break
			}
		}
		for i := range titles {
			fmt.Printf("%s (%s)\n", titles[i], urls[i])
		}
	}
}
