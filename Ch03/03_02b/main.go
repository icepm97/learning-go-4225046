package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	var i int
	var p *int
	fmt.Println(i)
	// pointer. currently nil
	fmt.Println(p)
	// dereferencing the pointer. panic (nil pointer dereference).
	// fmt.Println(*p)

	fmt.Println()
	if p == nil {
		fmt.Println("nil pointer")
	} else {
		fmt.Println(*p)
	}

	fmt.Println()
	intValue := 44
	intPointer := &intValue
	fmt.Printf("Pointer Type: %T\n", intPointer)
	fmt.Printf("Pointer: %v\n", intPointer)
	fmt.Printf("Pointer Value: %v\n", *intPointer)

	*intPointer = *intPointer / 2
	// when value referenced by intPointer is changed, intValue is also get changed.
	fmt.Printf("New Value: %v\n", intValue)
}
