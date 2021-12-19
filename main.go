package main

import (
	"fmt"
	"net/http"
)

func main() {
	var results = map[string]bool{}
	urls := []string{"https://www.google.com",
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
	for _, url := range urls {
		checkURL(url, results)
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func checkURL(url string, results map[string]bool) {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		results[url] = false
	}
	results[url] = true
}
