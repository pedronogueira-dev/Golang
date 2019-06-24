package main

import (
	"fmt"
	"net/http"
)

//Website Status Checker

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.co.uk",
	}

	c := make(chan string)

	for _, link := range links {
		go statusChecker(link, c)
	}

	for {
		go statusChecker(<-c, c)
	}
}

func statusChecker(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error Accessing", link, "might be down")
	} else {
		fmt.Println(link + " is Up")
	}

	c <- link
}
