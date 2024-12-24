package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey! What's your name?")
	
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
}
