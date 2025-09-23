package main

import "fmt"

func main() {
	fmt.Println("Math")

	i1, i2 := 1, 5
	// / is integer division for integers
	fmt.Println(i1 / i2)

	f1, f2 := 1.0, 5.0
	fmt.Println(f1 / f2)

	// opperands need to be in same type. can't do 1 / 1.0
	// compiler error invalid operation
	// fmt.Println(f1 / i2)

	// use type conversion
	fmt.Println(f1 / float64(i2))

	var f3 float32 = 1.0
	var f4 float64 = 2.0
	// types need to be same
	// fmt.Println(f3 + f4)
	fmt.Println(float32(f4) / f3)
	fmt.Println(1.5 / 2)
}
