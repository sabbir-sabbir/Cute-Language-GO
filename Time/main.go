package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now()
    formatteddTime := currentTime.Format("2006-01-02, Wednesday")
    fmt.Println(formatteddTime)
	fmt.Println("Welcome to 2025")
}
