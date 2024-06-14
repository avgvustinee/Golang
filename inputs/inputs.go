package main

import "fmt"

func main() {
	var temperature float32
	var sunny bool

	fmt.Println("Enter the current temperature : ")
	fmt.Scanf("%g", &temperature)
	fmt.Println("is the day sunny : ")
	fmt.Scanf("%t", &sunny)
	fmt.Printf("Current Temperature : %g\nis the day Sunny ? %t ", temperature, sunny)
}
