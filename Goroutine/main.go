package main

import (
	"fmt"
	"time"
)



func sayHello() {
	fmt.Println("HELLO- WORLD!")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("ok after 2 second")
}
func hi() {
	fmt.Println("HI HI HI")
}

func main() {
	fmt.Println("Learning... goroutine now ")
	go sayHello()
	hi()
	

	time.Sleep(3000 * time.Millisecond)
}
