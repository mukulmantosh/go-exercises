package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	var links []string
	links = append(links, "https://google.com")
	links = append(links, "https://facebook.com")
	links = append(links, "https://stackoverflow.com")
	links = append(links, "https://golang.org")
	links = append(links, "https://amazon.com")

	c := make(chan string) // Channel

	for _, link := range links {
		go checkLink(link, c)
	}

	//for {
	//	//fmt.Println(<-c)
	//	go checkLink(<-c, c)
	//}

	// Alternative Syntax
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	fmt.Println("Main goroutine finished")

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down !")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, " is up!")
	//c <- "It's Up !"
	c <- link
}
