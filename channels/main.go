package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://facebook.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		go func(link string){
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}

	fmt.Println(link, "is Up!")
	channel <- link
}
