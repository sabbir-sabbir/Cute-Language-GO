package main

import "fmt"


func Divide(a, b float64) (float64, error) {
      if b == 0 {
		return 0, fmt.Errorf("this error because the second argument of this func is 0")
	  }
	  return a / b, nil
}
func main() {
	fmt.Println("Starting")
	ans, err := Divide(5, 8)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(ans)
}