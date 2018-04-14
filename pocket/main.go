package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get access token
	f, err := os.Open("auth.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	a := Auth{}
	if err := json.Unmarshal(bs, &a); err != nil {
		panic(err)
	}

	handleJosuiWritings(a)
	fmt.Println("Saved all posts from blog http://blog.josui.me")
}
