package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
  res, err	:=  http.Get("https://jsonplaceholder.typicode.com/users")
  if err != nil {
	fmt.Println("Error getting", res.StatusCode)
	return
  }
  defer res.Body.Close()

//   Read data

  data, err := io.ReadAll(res.Body)
  if err!= nil {
    fmt.Println("Error reading body", err)
	return
 }
  fmt.Println(string(data))

}