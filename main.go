package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	var results = map[string]string{}
	urls := [13]string{"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.baidu.com",
		"https://www.wikipedia.org",
		"https://www.yandex.ru",
		"https://www.yahoo.com",
		"https://www.amazon.com",
		"https://www.netflix.com",
		"https://www.live.com",
		"https://www.zoom.us"}
	for i, url := range urls {
		fmt.Println("Waiting for check...", i)
		// channel을 만나면 loop에서는 건너뜀?
		go checkURL(url, c)
		fmt.Println("Next to channel", i)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func checkURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
