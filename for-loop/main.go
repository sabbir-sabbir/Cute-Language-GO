package main

import "fmt"

func main() {
	// for i:=0; i < 10; i++ {
    //       fmt.Println("This is a test")
	// }

	// counter := 0
	// for {
	// 	fmt.Println("This is a test of infinite loop")
	// 	counter++
	// 	if counter == 10 {
    //         break
    //     }
	// }

	// numbers :=[]int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	data := "SABBIR,  RAHMAN"
	for index, char := range data {
		fmt.Printf("index of data is %d and char of the data is %c ", index, char)
	}
}