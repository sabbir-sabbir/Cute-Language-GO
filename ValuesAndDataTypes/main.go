package main

import (
	"fmt"
	"strings"
)
func main() {

	str1 := "sabbir"
	str2 := "rahman"
	// result := strings.Join([]string{str1, str2}, " + ")
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println(result)


}