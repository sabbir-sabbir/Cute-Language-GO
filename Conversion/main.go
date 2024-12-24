package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 50
	str := strconv.Itoa(num)

	fmt.Println("String: ", str)
	fmt.Printf("Type: %T\n", str)

}