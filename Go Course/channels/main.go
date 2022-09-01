package main

import (
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://microsoft.com",
		"http://golang.org",
		"http://bitsaude.com.br",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link
		println("There is some problem in " + link)
		return
	}

	c <- link
	println(link + " is up!")
}
