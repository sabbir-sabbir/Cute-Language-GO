package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("testing")
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("Error getting", err)
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err!= nil {
        fmt.Println("Error reading body", err)
		return
    }
	fmt.Println(string(data))
}