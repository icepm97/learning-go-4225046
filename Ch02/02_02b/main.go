package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	number := 40

	stringLenght, err := fmt.Println(str1, str2, str3)
	if err == nil {
		fmt.Println("String length:", stringLenght)
	}
	
	fmt.Printf("Value of the number: %v\n", number)
	fmt.Printf("Type of the number: %T\n", number)

}
