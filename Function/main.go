package main

import (

  "fmt"

) 

//  func simpleFunction() {
// 	fmt.Println("This is a simple function")
//  }

// func add(a, b int) (int) {
// 	return a + b
// }

 func multiply(a, b int) (result int) {
 	result = a * b
	return
 }

func main() {
	// fmt.Println("Hello, World!")
	// simpleFunction()
	// ans := add(5, 10)
	// fmt.Println("add of two numbers:", ans)
	ans := multiply(5, 10)
	fmt.Println("multiply of two numbers:", ans)
}