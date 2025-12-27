package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://go.dev",
		"https://google.com",
		"https://duckduckgo.com",
		"https://localhost:8080",
		"https://archlinux.org/",
	}

	for _, link := range links {
		go checkLinks(link)
	}
}

func checkLinks(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link, "is working")
	}
}
