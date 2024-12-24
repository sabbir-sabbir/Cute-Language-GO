package main

import "fmt"

func main() {
	num := 5

	ptr := &num 

	fmt.Println("Value of num: ", num)
	fmt.Println("Value of num: ", ptr)
}