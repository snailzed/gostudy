package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	getUrlContent()
	return
	url := "https://www.baidu.com"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("get url error.")
	}
	fmt.Printf("%v", res)
}

func getUrlContent() {
	var url string
	flag.StringVar(&url, "u", "", "url: fetch url")
	flag.Parse()
	if url == "" {
		fmt.Println("url is empty.")
		return
	}
	if 0 != strings.Index(url, "http") {
		fmt.Println("url is invalid.")
		return
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("fetch url error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("Fetch content:", string(body))
}
