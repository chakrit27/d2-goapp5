package main

import "fmt"

func main() {
	var value interface{}
	value = "welcome"
	show(value)

	value = 49
	show(value)
}

func show(value interface{}) {
	fmt.Printf("%v, %T\n", value, value)
}