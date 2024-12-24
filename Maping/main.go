package main

import "fmt"

func main() {

	// grades of studens
	studensGrades := make(map[string]int)  

	// adding grades of studens
	studensGrades["Sabbir"] = 100
	studensGrades["Rahim"] = 90
	studensGrades["Karim"] = 80
	studensGrades["Jabir"] = 70
	studensGrades["Kabir"] = 60
     
	// print grades of studens
	// fmt.Println("Grades of Sabbir:", studensGrades["Sabbir"])
	// studensGrades["Sabbir"] = 700
	// fmt.Println(" new Grades of Sabbir:", studensGrades["Sabbir"])
	
	//  value, exits :=   studensGrades["Sabbir"]
	//  fmt.Println("Grades of Sabbir:", value, "Exits:", exits)

	for index, value := range studensGrades{
		fmt.Println("Index:", index, "Value:", value)
	}

}