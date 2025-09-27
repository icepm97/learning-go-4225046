package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")
	value := 11
	var result string
	if value > 0 {
		result = "greater than zero"
	} else if value < 0 {
		result = "less than zero"
	} else {
		result = "zero"
	}
	fmt.Println(result)

	// we can define variables inside if conditin.
	// they are only available in the scope of if clause
	if value1 := 0; value1 < 0 {
		fmt.Println("below zero")
	}
	// can't access value1 outside the if clause
	// give error
	// fmt.Println(value1)
}
