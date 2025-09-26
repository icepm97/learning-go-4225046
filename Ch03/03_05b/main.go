package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	stateCodes := make(map[string]string)
	stateCodes["WA"] = "Washington"
	stateCodes["OR"] = "Oregon"
	stateCodes["CA"] = "Californi"
	fmt.Println(stateCodes)

	fmt.Println(stateCodes["CA"])
	stateCodes["CA"] = "California"
	fmt.Println(stateCodes)
	delete(stateCodes, "OR")
	fmt.Println(stateCodes)

	// for loop
	fmt.Println()
	for k, v := range stateCodes {
		fmt.Printf("%v: %v\n", k, v)
	}

	// sorted for loop
	fmt.Println()
	keys := make([]string, len(stateCodes))
	i := 0
	for k := range stateCodes {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i := range keys {
		fmt.Printf("%v: %v\n", keys[i], stateCodes[keys[i]])
	}
}
