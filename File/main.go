package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a file
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println("Something went wrong", err)
	// 	return
	// } 
	// defer file.Close()

	// content := "Hello from Go!"
	// _, errors := io.WriteString(file, content+"\n")
	// if errors != nil {
	// 	fmt.Println("Something went wrong", errors)
	// 	return
	// }
	// fmt.Println("File written successfully")

	// Open a file
	file, err := os.Open("test.txt") 
	if err != nil {
		fmt.Println("File is not opening", err)
		return
	}
	defer file.Close()

	// make a buffer
	buffer := make([]byte, 1024)

	// Read the file
	for {
		num, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err !=  nil {
			fmt.Println("error reading the file", err)
			return
		}

		// Print the content of the file
		fmt.Println(string(buffer[:num]))

	}
}