package main

import "fmt"

type person struct {
	Fname string
	Lname string
	Age   int
}

func main() {
	var sabbir person
	sabbir.Fname = "Sabbirrrr"
	sabbir.Lname = "Rahmnan"
	sabbir.Age = 25
	
	fmt.Println(sabbir)

}