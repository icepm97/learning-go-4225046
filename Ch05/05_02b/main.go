package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	// fmt.Printf("The %v says %v!\n", dog.Breed, dog.Sound)
	dog.Speak()
	dog.Sound = "Arf"
	println(dog.SpeakThreeTimes())
}

type Dog struct {
	Bread string
	Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Bread, d.Sound)
}

func (d Dog) SpeakThreeTimes() string {
	return fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
}

// no overloading. error: already defined
// func (d Dog) Speak(i int) {}
