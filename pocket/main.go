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

type Auth struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

var auth = func() Auth {
	f, err := os.Open("auth.json")
	handleError(err)
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	handleError(err)

	a := Auth{}
	if err := json.Unmarshal(bs, &a); err != nil {
		panic(err)
	}
	return a
}()

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Pocket struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

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

func (p Pocket) Add(url string) { // rate limit: 320 times/hour
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

func saveToPocket(url string) {
	body := []byte(fmt.Sprintf(`{
		"url": "%s",
		"consumer_key": "%s",
		"access_token": "%s"
	}`, url, auth.ConsumerKey, auth.AccessToken))
	req, err := http.Post("https://getpocket.com/v3/add", "application/json", bytes.NewReader(body))
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + " fail to save the article to pocket whose url is: " + url)
	}
}

type action struct {
	Action string `json:"action"`
	URL    string `json:"url"`
}

func saveMultipleToPocket(urls []string) {
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
		ConsumerKey: auth.ConsumerKey,
		AccessToken: auth.AccessToken,
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

// Usage: go run *.go
func main() {
	/*
		handleJosuiWritings()
		fmt.Println("Saved all posts from blog http://blog.josui.me")

		handleYinWang()
		fmt.Println("Saved all posts from blog http://www.yinwang.org/")

		handleYinWangLofter()
		fmt.Println("Saved all posts from blog http://yinwang0.lofter.com/")

		handleLeetcodeArticle()
		fmt.Println("Saved all posts from site https://leetcode.com/articles/")

		handleMiaoHu()
		fmt.Println("Saved all posts from blog https://miao.hu/")

		handleLepture()
		fmt.Println("Saved all posts from blog https://lepture.com/")
	*/

	handleLiQi()
	fmt.Println("Saved all posts from blog http://liqi.io/")
}
