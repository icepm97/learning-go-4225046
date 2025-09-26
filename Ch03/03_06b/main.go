package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	terry := Dog{}
	terry.Breed = "Terry"
	fmt.Println(terry)
	fmt.Printf("%+v\n", terry)

	poodle := Dog{"Poodle", 45}
	fmt.Printf("%+v\n", poodle)
	fmt.Println(poodle.Weight)

}

type Dog struct {
	Breed  string
	Weight int
}
