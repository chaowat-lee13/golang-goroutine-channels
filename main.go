package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://go.dev",
		"http://google.co.th",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// infinite loop
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}