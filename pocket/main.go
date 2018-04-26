package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// read the article about how to get access token: http://www.cnblogs.com/febwave/p/4242333.html

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Pocket holds some keys
type Pocket struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// NewPocket creates a Pocket structure for operations
func NewPocket() Pocket {
	f, err := os.Open("auth.json")
	handleError(err)
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	handleError(err)

	p := Pocket{}
	if err := json.Unmarshal(bs, &p); err != nil {
		panic(err)
	}
	return p
}

// Add adds a url to pocket
// rate limit: 320 times/hour
func (p Pocket) Add(url string) {
	body := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		URL         string `json:"url"`
	}{
		ConsumerKey: p.ConsumerKey,
		AccessToken: p.AccessToken,
		URL:         url,
	}
	bs, err := json.Marshal(body)
	handleError(err)
	req, err := http.Post("https://getpocket.com/v3/add", "application/json", bytes.NewReader(bs))
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + " fail to save the article to pocket whose url is: " + url)
	}
}

type action struct {
	Action string `json:"action"`
	URL    string `json:"url"`
}

// AddMultiple adds multiple urls at one time
func (p Pocket) AddMultiple(urls []string) {
	actions := []action{}
	for _, url := range urls {
		actions = append(actions, action{
			Action: "add",
			URL:    url,
		})
	}
	body := struct {
		ConsumerKey string   `json:"consumer_key"`
		AccessToken string   `json:"access_token"`
		Actions     []action `json:"actions"`
	}{
		ConsumerKey: p.ConsumerKey,
		AccessToken: p.AccessToken,
		Actions:     actions,
	}
	bs, err := json.Marshal(body)
	handleError(err)
	req, err := http.Post("https://getpocket.com/v3/send", "application/json", bytes.NewReader(bs))
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + " fail to save articles: " + strings.Join(urls, "\n"))
	}
}

func (p Pocket) AddFake(urls []string) {
	fmt.Printf("[FAKE] Haved added %d articles\n", len(urls))
}

// Info stores some basic info for one site
type Info struct {
	URL       string                                    `json:"url"`
	URLSuffix string                                    `json:"url_suffix"`
	ListPath  string                                    `json:"list_path"`
	NextPath  string                                    `json:"next_path"`
	Skip      bool                                      `json:"skip"`
	Handler   func(*goquery.Selection) (string, string) `json:"handler"`
}

var sites = []Info{
	{
		URL:      "http://blog.josui.me",
		ListPath: "div.blog div.content article",
		NextPath: "nav.pagination a.pagination-next",
		Skip:     true,
		Handler:  handleJosuiWritings,
	},
	{
		URL:      "http://www.yinwang.org",
		ListPath: "li.list-group-item.title",
		Skip:     true,
		Handler:  handleYinWang,
	},
	/*
		handleYinWangLofter(p)
		handleLeetcodeArticle(p)
		handleMiaoHu(p)
		handleLepture(p)
		handleLiQi(p)
	*/
	{
		URL:      "http://jannerchang.bitcron.com",
		ListPath: "div.post_in_list",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		Skip:     true,
		Handler:  handleJannerChang,
	},
	{
		URL:      "https://blog.todoist.com",
		ListPath: "article.tdb-article-slat",
		NextPath: "div.tdb-pagination-holder a.next.page-numbers.nav__action",
		Skip:     true,
		Handler:  handleTodoist,
	},
	{
		URL:       "https://mymorningroutine.com",
		URLSuffix: "/routines/all/#continue-routine",
		ListPath:  "div#js-archive-list div.card-img.card-img--archive",
		Skip:      true,
		Handler:   handleMyMorningRoutine,
	},
	{
		URL:      "https://ulyssesapp.com/blog",
		ListPath: "main#main article[id]",
		NextPath: "nav.navigation.paging-navigation div.nav-previous a",
		Skip:     true,
		Handler:  handleUlysses,
	},
	{
		URL:       "https://www.scotthyoung.com/blog",
		URLSuffix: "/articles",
		ListPath:  "div#date-block ul li",
		Skip:      true,
		Handler:   handleScottHYoung,
	},
	{
		URL:      "https://startupnextdoor.com",
		ListPath: "main#content article[class]",
		NextPath: "nav.pagination a.older-posts",
		Skip:     false,
		Handler:  handleStartUpNextDoor,
	},
}

// Usage: go run *.go
func main() {
	p := NewPocket()

	for _, site := range sites {
		if site.Skip {
			fmt.Println("Skipped:", site.URL)
			continue
		}
		fmt.Println()
		fmt.Println("Started:", site.URL)
		url := site.URL + site.URLSuffix
		for {
			doc, err := goquery.NewDocument(url)
			handleError(err)

			list := doc.Find(site.ListPath)
			titles, urls := []string{}, []string{}
			list.Each(func(i int, s *goquery.Selection) {
				title, post := site.Handler(s)
				titles = append(titles, title)
				if strings.HasPrefix(post, site.URL) {
					urls = append(urls, post)
				} else {
					urls = append(urls, site.URL+post)
				}
			})

			p.AddMultiple(urls)
			// p.AddFake(urls)
			fmt.Printf("Saved %d articles from site %s to Pocket\n", len(titles), url)
			for i := range titles {
				fmt.Printf("%d. %s (%s)\n", i+1, titles[i], urls[i])
			}

			if site.NextPath == "" {
				break
			}
			next := doc.Find(site.NextPath)
			nextURL, exist := next.Attr("href")
			if !exist {
				break
			}
			if strings.HasPrefix(nextURL, site.URL) {
				url = nextURL
				continue
			}
			url = site.URL + nextURL
		}
		fmt.Println("Finished:", site.URL)
	}
}
