package main

import (
	"fmt"
	"strconv"
)

//Defining type struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Method (value reciver)
func (p Person)hello()string {
	return p.firstName + " " + p.lastName + " " + p.city + p.gender + " " + strconv.Itoa(( p.age))
}

func main() {
	//initialization
	p1 := Person{firstName: "Mark", lastName: "Saksberg", city: "Trat", gender: "Male", age: 25}
	fmt.Println(p1.hello())
}

