package main

import "fmt"

func main() {
	// Switch Optional Statement with Multiple cases
	switch daysOfWeek := "Saturday"; daysOfWeek {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
		fallthrough
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")

	}

}
