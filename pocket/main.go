package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

// Usage: go run *.go
func main() {
	/*
		handleJosuiWritings()
		fmt.Println("Saved all posts from blog http://blog.josui.me")

		handleYinWang()
		fmt.Println("Saved all posts from blog http://www.yinwang.org/")

		handleYinWangLofter()
		fmt.Println("Saved all posts from blog http://yinwang0.lofter.com/")
	*/

	handleLeetcodeArticle()
}
