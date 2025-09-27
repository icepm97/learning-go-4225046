package main

import (
)

func main() {
	do()
}

func do() {
	println("Do")
	s1 := addTwoValues(1, 5)
	println(s1)
	c1, s2 := addAllValues(1, 2, 3)
	println(c1)
	println(s2)
}

func addTwoValues(v1, v2 int) int {
	return v1 + v2
}

func addAllValues(vs ...int) (int, int) {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return len(vs), sum
}
