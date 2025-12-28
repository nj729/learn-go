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
	c := make(chan string)
	for _, link := range links {
		go checkLinks(link, c)
	}
	fmt.Println(len(links))
	for i := 0; i < len(links)-1; i++ {
		fmt.Println(<-c)
	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		//c <- "Site is down"
	} else {
		fmt.Println(link, "is working")
		c <- "yep its working"
	}
}
