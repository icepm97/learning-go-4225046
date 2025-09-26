package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Black"
	fmt.Println(colors)
	fmt.Printf("Type: %T\n", colors)
	fmt.Printf("Length: %v\n", len(colors))
	fmt.Println(colors[0])
	fmt.Printf("Uninitialized Value: %v, Type: %T\n", colors[2], colors[2])
	// uninitialized value is an empty string
	fmt.Println(colors[2] == "")

	fmt.Println()
	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

}
