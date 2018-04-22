package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

// Info stores some basic info for one site
type Info struct {
	URL     string               `json:"url"`
	Skip    bool                 `json:"skip"`
	Handler func(Pocket, string) `json:"handler"`
}

var sites = []Info{
	{
		URL:     "http://blog.josui.me",
		Skip:    true,
		Handler: handleJosuiWritings,
	},
	{
		URL:     "http://www.yinwang.org",
		Skip:    true,
		Handler: handleYinWang, // TODO: update
	},
	{
		URL:     "https://blog.todoist.com",
		Skip:    true,
		Handler: handleTodoist, // TODO: update
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
		fmt.Println("Started:", site.URL)
		site.Handler(p, site.URL)
		fmt.Println("Finished:", site.URL)
	}
	/*
		handleYinWangLofter(p)

		handleLeetcodeArticle(p)

		handleMiaoHu(p)

		handleLepture(p)

		handleLiQi(p)

		handleJannerChang(p)
	*/
}
