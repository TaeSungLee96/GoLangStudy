package main

import "fmt"

func main() {
	// slice is not limited to the length
	// array is limited to the length
	// []is a slice
	// [...]is a array
	names := []string{"a", "b", "c", "d", "e"}
	names[4] = "lala"               // change the value
	names = append(names, "Teamin") // add the value

	fmt.Println(names)
}
