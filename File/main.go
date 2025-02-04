package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//   First create a file...
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
        return
	}
	defer file.Close()

	content := "Hello this is weritten by Go programming Language"
	_, errors := io.WriteString(file, content+"\n")
	if errors != nil {
		fmt.Println("Error writing to file:", errors)
        return
	}
	fmt.Println("String Writing sucksexFully")
}