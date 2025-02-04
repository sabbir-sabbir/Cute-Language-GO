package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var num = 50
	// str := strconv.Itoa(num)

	// fmt.Println("String: ", str)
	// fmt.Printf("Type: %T\n", str)

	number_str := "12345"

	str_num, _ := strconv.Atoi(number_str)
	fmt.Println(str_num)
	fmt.Printf("Type of this value %T", str_num)

}