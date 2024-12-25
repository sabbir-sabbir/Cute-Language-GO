package main

import (
	"fmt"
	"time"
)





func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Say hello func ended successfully")
}
func sayHi() {
	fmt.Println("HI")
}

func main() {
	fmt.Println("Learning Goroutine")
	sayHello()
	sayHi()

	// waiting for first func execute
	time.Sleep(3000 * time.Millisecond)
}