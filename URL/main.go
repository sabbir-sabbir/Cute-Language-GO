package main

import (
	"fmt"
	"net/url"
)

func main() {
	
	myUrl := "https://chatgpt.com/c/67694f16-51e0-800f-bd60-9af8d94fad1d"
	parsedURL, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("URL Parse Error:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("hello world")
}
