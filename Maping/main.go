package main

import "fmt"

func main() {
    // myData := make(map[string]int)
	// myData["John"] = 30
	// myData["Jane"] = 25
	// myData["Bob"] = 35
	// myData["Alice"] = 28
	// myData["Tom"] = 40
	// fmt.Println("Marks of Alice:", myData["Alice"])

	// grades, exists := myData["sabbir"]
	// fmt.Println(grades,exists)

	person := map[string]int{
		"John":   30,
        "Jane":   25,
        "Bob":    35,
        "Alice":   28,
        "Tom":    40,

	}
	
	for index, val := range person {
		fmt.Printf("Index: %s\n Value: %d\n", index, val)
	}

}