package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")

	fileName := "./textFile.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkErrors(err)
	length, err := io.WriteString(file, "Hello from Go!")
	checkErrors(err)
	fmt.Printf("Wrote %v bytes.", length)

	println("Reading the file")
	readFile(fileName)
}

func checkErrors(err error){
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkErrors(err)
	fmt.Printf("%v\n", data)
	println(string(data))
}
