package main

import (
	"fmt"
	"io"
	"os"

)

func main() {
	// //   First create a file...
	// file, err := os.Create("index.html")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
    //     return
	// }
	// defer file.Close()
   
	// // add first content
	// content := "<h1>Hello this is a paragraph</h>"
	// _, errors := io.WriteString(file, content+"\n")
	// if errors != nil {
	// 	fmt.Println("Error writing to file:", errors)
    //     return
	// }
	// fmt.Println("String Writing sucksexFully")

       	// Open the file in append mode
	file, err := os.OpenFile("index.html", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// New content to append
	newContent := "Abad pukur bazar"

	// Write to the file
	_, errors := io.WriteString(file, newContent+"\n")
	if errors != nil {
		fmt.Println("Error writing to file:", errors)
		return
	}

	fmt.Println("String written successfully")


}