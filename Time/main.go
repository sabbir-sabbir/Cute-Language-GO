package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	// formattedTime := currentTime.Format("02-01-2006,  Monday, 3:04 PM" )
	// fmt.Println(formattedTime)


	newDate := currentTime.Add(48 * time.Hour)
	// fmt.Println(newDate)
	formattedNewDate := newDate.Format("02-01-2006,  Monday, 3:04 PM" )
	fmt.Println(formattedNewDate)

}