package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// put the file into a folder (e.g.: ~/CodeNewbie), and then execute `go run main.go`
func main() {
	url := "https://www.codenewbie.org/podcast/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	articles := doc.Find("article")
	articles.Each(func(i int, s *goquery.Selection) {
		as := s.Find("a")
		title := as.Eq(1).Find("h3").Text()
		fmt.Println(title)
		path, exists := as.Eq(0).Attr("href")
		if !exists {
			panic("href do not exist")
		}
		part := strings.Split(path, "/")[2]
		newURL := url + part
		newDoc, err := goquery.NewDocument(newURL)
		if err != nil {
			panic(err)
		}
		src, exists := newDoc.Find("source").Attr("src")
		if !exists {
			fmt.Println("no src value for url: ", newURL)
			return
		}
		if !strings.HasSuffix(src, "mp3") {
			fmt.Println("no mp3 file in src: ", src)
			return
		}
		fmt.Println(src)
		// if a file with the same name exists, check its size
		tf, err := os.Open(title + ".mp3")
		if err == nil {
			tfInfo, err := tf.Stat()
			if err != nil {
				panic(err)
			}
			if tfInfo.Size() > 1*1024*1024 { // if the size of one file is more than 1M, then it is supposed to be downloaded correctly
				return
			}
		}
		defer tf.Close()
		// if not exist
		fmt.Println("Create a new file: " + title + ".mp3")
		out, err := os.Create(title + ".mp3")
		if err != nil {
			panic(err)
		}
		defer out.Close()
		fmt.Println("Download mp3 file")
		resp, err := http.Get(src)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("Copy content from response to file")
		if _, err := io.Copy(out, resp.Body); err != nil {
			panic(err)
		}

		// Why do I need it?
		// After downloading for a while, it may not work properly; that is, you will get some files whose sizes are smaller than 1M
		// So I ask it to sleep 1 minute before downloading a new episode
		// But it seems that 1 minute is not long enough
		// If you find that it occurs, quit and re-run it
		fmt.Println("Sleep 1 minute")
		time.Sleep(time.Minute)
		fmt.Println("起床搬砖啦")
	})
}
