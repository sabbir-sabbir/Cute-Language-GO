package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!, Now we are trying to arrays")
	// var name[5] string
	// name[0] = "Hello"
	// name[1] = "World"
	// fmt.Println(name[0], name[1])

	// name := [3]string{"Hello", "World", "Go"}
	// fmt.Println(name[0], name[1], name[2])

	var name  [5]string

	name[0] = "Sabbir"
	name[4] = "Rahman"

	fmt.Printf("Show the array: %q\n", name)
}