package main

import "fmt"

func main() {
 x := 2
 y := 2

//  if x > 10 {
// 	fmt.Println("x is greater than 10")
//  } else if x < 10 {
// 	fmt.Println("x is less than 10")
//  } else  {
// 	fmt.Println("x is equal to 10")
//  }

if x > 5 && y > 5 {
	fmt.Println("x and y are greater than 5")
} else {
	fmt.Println("we are learning go")
}

}