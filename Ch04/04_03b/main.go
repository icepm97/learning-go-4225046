package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Printf("%v: %v\n", i, colors[i])
	}

	println()
	for i := range colors{
		fmt.Printf("%v: %v\n", i, colors[i])
	}

	println()
	for i, color := range colors {
		fmt.Printf("%v: %v\n", i, color)
	}

	println()
	states := make(map[string]string)
	states["R"] = "RED"
	states["G"] = "GREEN"
	states["B"] = "BLUE"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	
	println()
	for k := range states {
		fmt.Printf("%v: %v\n", k, states[k])
	}

	println()
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	// continue & break are also supported

	println()
	i = 1
	for true {
		if i > 200 {
			goto theEnd
		}
		i += i
	}
	theEnd: println(i)
}
