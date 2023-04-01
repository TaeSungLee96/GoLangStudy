// 1. Variables and Constants
package main

import "fmt"

func printVariable() {
	var name bool = false
	name = true
	fmt.Println(name)
}

func printConstant() {
	const name bool = false
	fmt.Println(name)
}

func printAbbreviation() {
	bool := false // bool is a keyword
	bool = true
	// bool = 1 // Error
	// bool = "string" // Error
	// bool = 1.0 // Error
	fmt.Println(bool)
}

func main() {
	printVariable()
	printConstant()
	printAbbreviation()
}
