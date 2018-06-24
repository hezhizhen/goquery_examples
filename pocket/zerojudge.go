package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func handleZeroJudge(url string, p Pocket) {
	var urls []string
	page := 0
	for {
		fmt.Println(url + fmt.Sprintf("/P%d", page))
		doc, err := goquery.NewDocument(url + fmt.Sprintf("/P%d", page))
		handleError(err)
		list := doc.Find("div.blogbody")
		list.Each(func(i int, s *goquery.Selection) {
			href, exist := s.Find("h3.title.brk_h a").Attr("href")
			if !exist {
				panic("missing url")
			}
			urls = append(urls, strings.TrimSuffix(url, "/zerojudge")+href)
		})
		if list.Length() < 10 {
			break
		}
		page++
	}
	p.AddMultiple(urls)
	for i := range urls {
		fmt.Printf("%3d: %s\n", i+1, urls[i])
	}
}
