package main

import (

  "fmt"

) 

func simpleFunc (){
	fmt.Println("its a simple func")
}

func add(a, b int) (int) {
	return a + b
}

func main() {
	fmt.Println("Hello world!")
	simpleFunc()
	add := add(1, 2)
	fmt.Println("This wll add two numbers: ", add)
} 