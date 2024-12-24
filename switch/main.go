package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1:
		fmt.Println("friday day of the week")
	case 2:
		fmt.Println("saturday is the second day of the week")
	case 3:
		fmt.Println("sunday is the third day of the week")
	case 4:	
		fmt.Println("monday is the fourth day of the week")
	case 5:
		fmt.Println("tuesday is the fifth day of the week")
	default:
		fmt.Println("not a valid day")
	}
}