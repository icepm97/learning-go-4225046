package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	go Say("Hello from go routine")
	go func() { println("Anonymous go routine") }()
	println("Hello from main")
	// program will die after completing main.
	// sleep to wait till go routine is complete
	time.Sleep(time.Second)
	println("All done")
}

func Say(s string) {
	time.Sleep(1 * time.Second)
	println(s)
}
