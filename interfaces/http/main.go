package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 9999)
	//resp.Body.Read(bs)
	//sb := string(bs)
	//log.Printf(sb)

	io.Copy(os.Stdout, resp.Body)
}
