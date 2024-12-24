package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int   `json:"age"`
    Single bool `json:"single"`
}

func main() {
	Theperson := Person{Name: "John", Age: 30, Single: true}
   jsonData, err :=  json.Marshal(Theperson)
   if err != nil {
    fmt.Println("Error in marshalling" , err)
    return
   }
//    fmt.Println(string(jsonData))

    var person Person 
    err = json.Unmarshal(jsonData, &person)
    if err != nil {
        fmt.Println("Error in unmarshalling" , err)
        return
    }
    fmt.Println(person)


}