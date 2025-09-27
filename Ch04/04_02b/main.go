package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	default:
		result = "Weekend"
	}
	fmt.Println(result)

	fmt.Println()
	x := 0
	// with fallthrough, switch act as a switch in c lang
	// otherwise go doesn't need break statement
	switch {
	case x < 0:
		fmt.Println("less than zero")
		fallthrough
	case x == 0:
		fmt.Println("zero")
		fallthrough
	default:
		fmt.Println("greater than zero")
	}
}
