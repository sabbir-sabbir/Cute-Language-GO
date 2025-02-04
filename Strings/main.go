package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	ran_str := "Hello Hello This Is The Go Language Can You Have Your Friends With Your"
	count := strings.Count(ran_str, "Hello")
	fmt.Println(count)
}