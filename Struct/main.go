package main

import "fmt"

type Person struct {
	Name string
	Age  int
    single bool
}
type Contact struct {
	Email string
	phone string
}

type Address struct {
	House int
	Aria string
	state string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {

  var Employe1 Employee

  Employe1.Person_Details = Person{
	Name: "John",
    Age:  30,
    single: true,
  }

  Employe1.Person_Contact = Contact{
    Email: "john@example.com",
	phone: "888 888 999",
 }

  Employe1.Person_Address = Address{
	House: 123,
    Aria:  "New York",
    state: "NY",
  }
  fmt.Println(Employe1)

}
  