package main

import "fmt"


func divite(a,b float64) (float64, error) {
   if b == 0 {
	return 0, fmt.Errorf("can't divide by zero")

   }

	return a/b, nil
}

func main() {
	fmt.Println("hello world, Now we are trying to handle errors")
	ans, _ := divite(10, 4)
	// if err != nil {
	// 	fmt.Println("error is detected")
	// }
	fmt.Println("Divition of two numbers:", ans)
}