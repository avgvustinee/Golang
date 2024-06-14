package main

import "fmt"

func main() {
	// explicitly declare the data type
	var number1 int = 10
	fmt.Println(number1)
	// assign a value without declaring the data type
	var number2 = 20
	fmt.Println(number2)
	// shorthand notation to define variable
	number3 := 30
	println(number3)
	// constants
	const mass = 23
	fmt.Println("The mass is :", mass)
	// printf()
	name := "Augustine"
	age := 24

	fmt.Printf("%s is %d years old .", name, age)
}
