package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")

	date := time.Date(2004, 12, 1, 5, 50, 0, 0, time.UTC)
	fmt.Printf("Old date: %s\n", date)

	now := time.Now()
	fmt.Printf("Now: %s\n", now)
	fmt.Printf("Type: %T\n", now)

	fmt.Println(now.Format(time.ANSIC))
	// time formatting is different.
	// to represent year go uses 2006 instead of yyyy
	// month - 1, date - 2, hour - 3, minutes - 4, seconds - 5, year - 6
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	// test usage of panic
	// execution stops at panic
	panic("error")
	fmt.Println("Print after error")
}
