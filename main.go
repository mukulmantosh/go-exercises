package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
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
	c <- "It's Up !"
}
