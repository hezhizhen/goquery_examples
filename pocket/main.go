package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Info stores some basic info for one site
type Info struct {
	URL       string                                    `json:"url"`
	URLSuffix string                                    `json:"url_suffix"`
	StartURL  string                                    `json:"start_url"`
	ListPath  string                                    `json:"list_path"`
	TitlePath string                                    `json:"title_path"`
	URLPath   string                                    `json:"url_path"`
	NextPath  string                                    `json:"next_path"`
	Skip      bool                                      `json:"skip"`
	Fake      bool                                      `json:"fake"`
	Handler   func(*goquery.Selection) (string, string) `json:"handler"`
}

var sites = []Info{
	{
		URL:       "https://blog.brickgao.com",
		ListPath:  "div.post-summary",
		TitlePath: "div.post-title a",
		URLPath:   "div.post-title a",
		NextPath:  "div.paginator a.extend.next",
		Skip:      true,
	},
	// now in descending order
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
		Skip:     true,
		Handler:  handleStartUpNextDoor,
	},
	{
		URL:      "https://blog.trello.com",
		ListPath: "div.post-item",
		NextPath: "div.blog-pagination a.next-posts-link",
		Skip:     true,
		Handler:  handleTrello,
	},
	{
		URL:      "http://www.catcoder.com",
		ListPath: "section#posts article",
		NextPath: "nav.pagination a.extend.next",
		Skip:     true,
		Handler:  handleCatCoder,
	},
	{
		URL:      "https://www.waerfa.com",
		ListPath: "main#main article[id]",
		NextPath: "nav.navigation.posts-navigation div.nav-previous a",
		Skip:     true,
		Handler:  handleWaerfa,
	},
	{
		URL:      "https://unee.wang",
		ListPath: "div.post-list div.mod-post",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		Skip:     true,
		Handler:  handleUneeWang,
	},
	{
		URL:      "https://xiaomu.bitcron.com",
		ListPath: "div.post",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		Skip:     true,
		Handler:  handleXiaomu,
	},
	{
		URL:      "https://deans.site",
		ListPath: "div.post_list div.post",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		Skip:     true,
		Handler:  handleDeansSite,
	},
	{
		URL:      "http://www.asianefficiency.com/blog",
		ListPath: "article[class]",
		NextPath: "nav.archive.pagination div.next a",
		Skip:     true,
		Handler:  handleAsianEfficiency,
	},
	{
		URL:      "https://usesthis.com",
		ListPath: "article.interviewee.h-card",
		NextPath: "nav#paginator a#next",
		Skip:     true,
		Handler:  handleUseThis,
	},
	{
		URL:      "https://blog.jez.io",
		ListPath: "article.hentry div.entry-wrapper",
		Skip:     true,
		Handler:  handleJez,
	},
	{
		URL:      "http://wsfdl.com",
		ListPath: "ul.post-list li",
		Skip:     true,
		Handler:  handleWsfdl,
	},
	{
		URL:      "https://maqmodo.com",
		ListPath: "div.blogpostcategory",
		NextPath: "div.wp-pagenavi a.nextpostslink",
		Skip:     true,
		Handler:  handleMaqmodo,
	},
	{
		URL:       "https://productivityist.com",
		URLSuffix: "/category/blog/",
		ListPath:  "article[class]",
		NextPath:  "li.pagination-next a",
		Skip:      true,
		Handler:   handleProductivityist,
	},
	{
		URL:      "https://www.appinn.com",
		ListPath: "div#spost div.spost.post",
		NextPath: "div.navigation a.nextpostslink",
		Skip:     true,
		Handler:  handleAppinn,
	},
	{
		URL:       "http://gank.io",
		URLSuffix: "/history",
		ListPath:  "li div.row",
		Skip:      true,
		Handler:   handleGankIO,
	},
	{
		URL:       "http://onespiece.strikingly.com",
		URLSuffix: "/blog/df60b9b6a7b",
		NextPath:  "span.s-blog-footer-btn.s-blog-footer-previous a",
		Skip:      true,
	},
	{
		URL:      "http://blog.yuelong.info",
		ListPath: "section#posts article[class]",
		NextPath: "div.alignleft a",
		Skip:     true,
		Handler:  handleYuelong,
	},
}

// read the article about how to get access token:
// http://www.cnblogs.com/febwave/p/4242333.html

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
	req, err := http.Post(
		"https://getpocket.com/v3/add",
		"application/json",
		bytes.NewReader(bs),
	)
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + "fail to save the article to pocket whose url is:" + url)
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
	req, err := http.Post(
		"https://getpocket.com/v3/send",
		"application/json",
		bytes.NewReader(bs),
	)
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + " fail to save articles: " + strings.Join(urls, "\n"))
	}
}

// AddFake is used to show number of articles
func (p Pocket) AddFake(urls []string) {
	fmt.Printf("[FAKE] Haved added %d articles\n", len(urls))
}

// Usage: go run *.go
func main() {
	p := NewPocket()

	for _, site := range sites {
		if site.Skip {
			fmt.Println("Skipped:", site.URL)
			continue
		}
		fmt.Println("Started:", site.URL)
		url := site.URL + site.URLSuffix
		total := 0
		if site.StartURL != "" && url != site.StartURL {
			url = site.StartURL
		}
		for {
			doc, err := goquery.NewDocument(url)
			handleError(err)

			titles, urls := []string{}, []string{}
			if site.ListPath != "" {
				list := doc.Find(site.ListPath)
				list.Each(func(i int, s *goquery.Selection) {
					var title, post string
					if site.Handler == nil {
						var exist bool
						post, exist = s.Find(site.URLPath).Attr("href")
						if !exist {
							panic("missing url")
						}
						title, exist = s.Find(site.TitlePath).Attr("title")
						if !exist {
							title = s.Find(site.TitlePath).Text()
						}
					} else {
						title, post = site.Handler(s)
					}
					titles = append(titles, title)
					if strings.HasPrefix(post, site.URL) {
						urls = append(urls, post)
					} else {
						urls = append(urls, site.URL+post)
					}
					total++
				})
			} else {
				urls = append(urls, url)
				titles = append(titles, doc.Find("head title").Text())
				total++
			}

			if site.Fake {
				p.AddFake(urls)
			} else {
				p.AddMultiple(urls)
				// time.Sleep(time.Second * 30) // avoid something
			}
			fmt.Printf("[%s] Saved %d articles from site %s to Pocket\n",
				time.Now().Format("2006-01-02 15:04:05"), len(titles), url)
			for i := range titles {
				fmt.Printf("                      %d. %s (%s)\n",
					i+1, titles[i], urls[i])
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
		fmt.Println("Finished:", site.URL, "( In all:", total, ")")
		fmt.Println()
	}
}
