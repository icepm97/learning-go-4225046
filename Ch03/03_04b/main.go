package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	// This is an slice. Notice: no length is specified.
	var colors1 = []string{"Red", "Green", "Blue"}
	fmt.Println(colors1)

	fmt.Println()
	var colors = make([]string, 0, 3)
	// index out of range error. No zero as of now
	// colors[0] = "Red"
	colors = append(colors, "Red", "Green")
	fmt.Println(colors)
	fmt.Println(len(colors))
	fmt.Println(colors[0])
	colors = append(colors, "Blue", "Purple", "Fuschia")
	fmt.Println(colors)

	// remove items from slice
	fmt.Println()
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	numbers = remove(numbers, 2)
	fmt.Println(numbers)

	// sort slice
	// inplace
	fmt.Println()
	sort.Strings(colors)
	fmt.Println(colors)

}

func remove(slice []int, i int) []int {
	// get values before i and append items after i using spread operator
	return append(slice[:i], slice[i+1:]...)
}
