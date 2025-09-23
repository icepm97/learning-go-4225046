package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter a value: ")
	// str, _ := reader.ReadString('\n')
	// fmt.Println(str)

	str, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	// parse float supports inf, -inf and nan
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err)
	}

}
