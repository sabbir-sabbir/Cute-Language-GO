package main

import "fmt"

func main() {
	fmt.Println("Starting the slice...")
	number := []int{1, 2, 3, 4,5,6}
	fmt.Println(number)
	number = append(number, 10,10,10,10,100,200,500,800,600,300)
	fmt.Println(number)
    fmt.Printf("Length of numbers now: %d\n", len(number))
	price := make([]int, 3, 5)
	price = append(price, 10,10,10)
    fmt.Println("slice", price)
	fmt.Println("Length", len(price))
	fmt.Println("Capacity", cap(price))
}