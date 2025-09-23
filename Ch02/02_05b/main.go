package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	radius := 5.0
	circumference := 2.0 * math.Pi * radius
	fmt.Printf("no rounding: %v\n", circumference)
	fmt.Printf("math.Round: %v\n", math.Round(circumference*100)/100)
	fmt.Printf("printf rounding: %0.2f\n", circumference)
}
