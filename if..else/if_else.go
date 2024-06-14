package main

import "fmt"

func main() {
	age := 15
	myAge := 24
	// if else example
	/*if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	} */

	// if else if example
	if age == myAge {
		fmt.Println("we are the same age")
	} else if age > myAge {
		fmt.Println("you are older than me")
	} else {
		fmt.Println("you are younger than me")
	}

}
