package main

import "fmt"

func main() {
	fmt.Println("This is Array")
	var name[5] string
	name[0] = "Hello"
	name[1] = "World"
	fmt.Println(name[0], name[1])

	number := [7]int{1, 2, 3, 4, 5}
	fmt.Println(number)
	fmt.Println(len(number))
	fmt.Println("_____________________________________________________")
	fmt.Println(name[1])
    price := [5]int{}
	fmt.Println(price)
	var numbers[50]string
	fmt.Println(numbers)
	fmt.Printf( "%q", numbers)
}