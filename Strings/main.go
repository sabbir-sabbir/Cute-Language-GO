package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := "one two three four five five "

	// count :=  strings.Count(str, "five")
	// fmt.Println("count", count)

	str1 := "sabbir"
	str2 := "rahman"
	result := strings.Join([]string{str1, str2}, " + ")

	fmt.Println(result)
}